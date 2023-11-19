package main

import (
	"log"
	"net"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	pb "github.com/zuno90/go-ws/pb"
	"github.com/zuno90/go-ws/routes"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedAuthServer
}

//
// func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
// 	log.Printf("Received: %v", in.GetName())
// 	return &pb.HelloReply{Resmessage: "Hello from server " + in.GetName()}, nil
// }

func main() {
	log.SetFlags(log.Lshortfile)

	env := os.Getenv("GO_ENV")
	if env != "production" {
		if err := godotenv.Load(); err != nil {
			log.Fatalf("Some error occured. Err: %s", err)
		}
	}
	// connect database
	// configs.ConnectPostgresDB()
	// go initGrpcServer() // init grpc server
	initFiberServer() // init websocket server
}

func initFiberServer() {
	app := fiber.New()

	// routes.SetUpHttpRoutes(app)
	routes.SetUpWebsocket(app)

	httpPort := os.Getenv("PORT")
	log.Fatal(app.Listen("localhost:" + httpPort))
	// Access the websocket server: ws://localhost:3000/ws/123?v=1.0
	// https://www.websocket.org/echo.html
}

func initGrpcServer() {
	grpcPort := os.Getenv("GRPC_PORT")
	listen, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("GRPC is listening on port %s", listen.Addr().String())
	grpcServer := grpc.NewServer()

	pb.RegisterAuthServer(grpcServer, &server{})

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve GRPC server %v", err)
	}
}
