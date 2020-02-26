package main

import (
	"context"
	"log"
	"time"

	pb "github.com/toyopilgrim/go-mascot-client-practice/grpc"
	"google.golang.org/grpc"
)

const (
	address = "localhost:6565"
)

func main() {
	log.Printf("Program started...")
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("could not connect : %v", err)
	}
	defer conn.Close()
	client := pb.NewMascotServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.Hello(ctx, &pb.MascotRequest{Id: 1})
	if err != nil {
		log.Fatalf("could not meet: %v", err)
	}
	log.Printf("Hey, it's %s ! %s", r.Name, r.Message)

}
