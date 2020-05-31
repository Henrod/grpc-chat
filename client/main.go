package main

import (
	"context"
	"io"
	"log"

	pb "github.com/Henrod/chat-example/protogen"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8000", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewFeedAPIClient(conn)

	ctx := context.Background()
	stream, err := client.Subscribe(ctx, &pb.SubscribeRequest{UserId: "1"})
	if err != nil {
		log.Fatal(err)
	}

	for {
		in, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		log.Print(in.Message)
	}
}
