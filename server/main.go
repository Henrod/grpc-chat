package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/Henrod/chat-example/protogen"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type FeedServer struct {
	feeds map[string]*Feed
}

func NewFeedServer() *FeedServer {
	return &FeedServer{
		feeds: map[string]*Feed{},
	}
}

type Feed struct {
	messages []*pb.Message
	subs     map[uuid.UUID]chan *pb.Message
}

func NewFeed() *Feed {
	return &Feed{
		subs: map[uuid.UUID]chan *pb.Message{},
	}
}

func (f *FeedServer) PostMessage(_ context.Context, r *pb.PostMessageRequest) (*pb.PostMessageResponse, error) {
	feed, ok := f.feeds[r.GetUserId()]
	if !ok {
		feed = NewFeed()
		f.feeds[r.GetUserId()] = feed
	}

	feed.messages = append(feed.messages, r.Message)
	for _, s := range feed.subs {
		s <- r.Message
	}

	return &pb.PostMessageResponse{Message: r.Message}, nil
}

func (f *FeedServer) ListMessages(_ context.Context, r *pb.ListMessagesRequest) (*pb.ListMessagesResponse, error) {
	feed, ok := f.feeds[r.GetUserId()]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "user not found: %v", r.GetUserId())
	}

	return &pb.ListMessagesResponse{Messages: feed.messages}, nil
}

func (f *FeedServer) Subscribe(r *pb.SubscribeRequest, s pb.FeedAPI_SubscribeServer) error {
	ctx := s.Context()

	c := make(chan *pb.Message, 1)
	cID := uuid.New()
	f.feeds[r.GetUserId()].subs[cID] = c
	for {
		select {
		case <-ctx.Done():
			delete(f.feeds[r.GetUserId()].subs, cID)
			log.Print("connection interrupted")
			return nil
		case msg := <-c:
			if err := s.Send(&pb.SubscribeResponse{
				Message: msg,
			}); err != nil {
				log.Printf("send subscribe message: %v", err)
				return status.Error(codes.Internal, "failed to send subscribe message")
			}
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterFeedAPIServer(s, NewFeedServer())

	log.Print("listening at localhost:8000")
	if err = s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
