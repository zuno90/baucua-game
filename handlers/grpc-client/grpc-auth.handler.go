package grpcclient

import (
	"context"
	"log"

	pb "github.com/zuno90/go-ws/pb"
)

func GetAuthUser(j string) (*pb.User, error) {
	conn, err := NewGrpcClient()
	if err != nil {
		log.Fatalf("Could not create connection grpc %v", err)
	}
	client := pb.NewAuthClient(conn)
	r, err := client.CheckAuth(context.Background(), &pb.Jwt{Jwt: j})
	if err != nil {
		log.Println("err", err)
		return nil, err
	}
	if r.GetId() == 0 {
		return nil, err
	}
	return r, nil
}
