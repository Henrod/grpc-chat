package main

import (
	"time"

	pb "github.com/Henrod/chat-example-2/protogen/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type FeedAPIV2 struct {
	feed map[string][]*pb.Message
}

func NewFeedAPIV2() *FeedAPIV2 {
	return &FeedAPIV2{feed: map[string][]*pb.Message{
		"1": {
			{Body: "msg1"},
			{Body: "msg2"},
			{Body: "msg3"},
		},
	}}
}

func (f *FeedAPIV2) ListMessages(r *pb.ListMessagesRequest, s pb.FeedAPI_ListMessagesServer) error {
	messages, ok := f.feed[r.GetFeedId()]
	if !ok {
		return status.Error(codes.NotFound, "user not found")
	}

	for _, msg := range messages {
		err := s.Send(&pb.ListMessagesResponse{Message: msg})
		if err != nil {
			return err
		}

		time.Sleep(3 * time.Second)
	}

	return nil
}
