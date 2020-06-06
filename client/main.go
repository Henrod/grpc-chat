package main

import (
	"context"
	"flag"
	"log"

	"google.golang.org/grpc/credentials"

	pb "github.com/Henrod/chat-example-2/protogen"
	"google.golang.org/grpc"
)

const certFile = "../creds/domain.crt"

func main() {
	creds, _ := credentials.NewClientTLSFromFile(certFile, "")
	dial, err := grpc.Dial("localhost:8000", grpc.WithTransportCredentials(creds))
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
