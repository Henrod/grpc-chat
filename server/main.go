package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"time"

	pb "github.com/Henrod/chat-example-2/protogen"
	"github.com/golang/protobuf/ptypes"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

type HenrodAPI struct{}

func (h *HenrodAPI) Echo(context.Context, *pb.EchoRequest) (*pb.EchoResponse, error) {
	return &pb.EchoResponse{}, nil
}

func startHTTPServer() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := pb.RegisterFeedAPIHandlerFromEndpoint(ctx, mux, "localhost:8000", opts); err != nil {
		return err
	}

	return http.ListenAndServe("localhost:8001", mux)
}

func main() {
	lis, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer(grpc.UnaryInterceptor(Metrics))
	pb.RegisterFeedAPIServer(s, NewFeedAPI())
	pb.RegisterHenrodAPIServer(s, &HenrodAPI{})

	go func() {
		log.Print("listening gRPC at localhost:8000")
		if err = s.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()

	log.Print("listening HTTP at localhost:8001")
	if err = startHTTPServer(); err != nil {
		log.Fatal(err)
	}
}
