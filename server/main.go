package main

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/golang/protobuf/ptypes"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/Henrod/chat-example-2/protogen"
	"google.golang.org/grpc"
)

type FeedAPI struct {
	feed map[string][]*pb.Message
}

func NewFeedAPI() *FeedAPI {
	return &FeedAPI{feed: map[string][]*pb.Message{}}
}

func (f *FeedAPI) PostMessage(
	_ context.Context,
	r *pb.PostMessageRequest,
) (*pb.PostMessageResponse, error) {
	messages, ok := f.feed[r.GetUserId()]
	if !ok {
		messages = f.feed[r.GetUserId()]
	}

	r.Message.CreatedAt = ptypes.TimestampNow()

	messages = append(messages, r.Message)
	f.feed[r.GetUserId()] = messages

	return &pb.PostMessageResponse{Message: r.Message}, nil
}

func (f *FeedAPI) ListMessages(
	_ context.Context,
	r *pb.ListMessagesRequest,
) (*pb.ListMessagesResponse, error) {
	messages, ok := f.feed[r.GetUserId()]
	if !ok {
		return nil, status.Error(codes.NotFound, "user not found")
	}

	time.Sleep(time.Second)

	return &pb.ListMessagesResponse{Messages: messages}, nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer(grpc.UnaryInterceptor(Metrics))
	pb.RegisterFeedAPIServer(s, NewFeedAPI())

	log.Print("listening at localhost:8000")
	if err = s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
