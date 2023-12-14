package utils

import (
	"context"
	"log"

	pb "github.com/zuno90/go-ws/pb"
	"google.golang.org/grpc"
)

func RunGRPCClientTest() {
	conn, err := grpc.Dial(":3030", grpc.WithBlock())
	if err != nil {
		log.Fatalf("could not connect %s", err)
	}
	defer conn.Close()
	c := pb.NewAuthClient(conn)
	r, err := c.CheckAuth(context.Background(), &pb.Jwt{Jwt: "zuno"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("User from GRPC server: '%s'", r)
}
