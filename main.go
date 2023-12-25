package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"github.com/zuno90/go-ws/configs"
	pb "github.com/zuno90/go-ws/pb"
	"github.com/zuno90/go-ws/routes"
	"github.com/zuno90/go-ws/utils"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedAuthServer
}

//	func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
//		log.Printf("Received: %v", in.GetName())
//		return &pb.HelloReply{Resmessage: "Hello from server " + in.GetName()}, nil
//	}

func main() {
	log.SetFlags(log.Lshortfile)
	loadConfigs() // load config

	go initGrpcServer()  // init grpc server
	go initFiberServer() // init websocket server

	// Wait for interrupt signal to gracefully shutdown
	waitForShutdownSignal()
}

func loadConfigs() {
	env := os.Getenv("GO_ENV")
	if env != "production" {
		if err := godotenv.Load(); err != nil {
			log.Fatalf("Some error occured. Err: %s", err)
		}
	}
	// Set up Viper to read environment variables
	viper.AutomaticEnv()
	// configs.ConnectPostgresDB() // connect database
	configs.ConnectKeydbServer() // connect keydb caching
}

func initFiberServer() {
	app := fiber.New()
	// routes.SetUpHttpRoutes(app)
	routes.SetUpWebsocket(app)

	httpPort := viper.GetString("PORT")
	log.Printf("Fiber server is listening on port haha %s", httpPort)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	log.Fatal(app.Listen(fmt.Sprintf("localhost:%s", httpPort)))

	// Access the websocket server: ws://localhost:3000/ws/123?v=1.0
	// https://www.websocket.org/echo.html
}

func initGrpcServer() {
	grpcPort := viper.GetString("GRPC_PORT")
	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("GRPC server is listening on port %s", listen.Addr().String())
	grpcServer := grpc.NewServer()

	pb.RegisterAuthServer(grpcServer, &server{})

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve GRPC server %v", err)
	}
}

func waitForShutdownSignal() {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	<-sigCh

	fmt.Println("Received shutdown signal. Shutting down gracefully...")
	// Graceful shutdown code goes here
	if err := utils.Clear(); err != nil {
		log.Fatal("Can not clear all keydb cache", err)
	}
	fmt.Println("Clear all keydb cache of logged users!")
	fmt.Println("Server shut down gracefully.")
}
