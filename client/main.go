package main

import (
	"context"
	"flag"
	"log"

	pb "github.com/Henrod/chat-example-2/protogen"
	"google.golang.org/grpc"
)

func main() {
	dial, err := grpc.Dial("localhost:8000", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	userID := flag.String("user_id", "", "user to get messages from")
	flag.Parse()

	client := pb.NewFeedAPIClient(dial)
	messages, err := client.ListMessages(context.Background(), &pb.ListMessagesRequest{
		UserId: *userID,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Print(messages)
}
