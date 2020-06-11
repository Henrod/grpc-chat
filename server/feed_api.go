package main

import (
	"context"

	"google.golang.org/genproto/protobuf/field_mask"

	pb "github.com/Henrod/chat-example-2/protogen"

	"github.com/golang/protobuf/ptypes"
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

	r.Message.CreatedTime = ptypes.TimestampNow()

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

	messages, err := f.fields(messages, r.Fields)
	if err != nil {
		return nil, err
	}

	return &pb.ListMessagesResponse{Messages: messages}, nil
}

func (f *FeedAPI) fields(messages []*pb.Message, mask *field_mask.FieldMask) ([]*pb.Message, error) {
	if mask == nil || len(mask.Paths) == 0 {
		mask = &field_mask.FieldMask{
			Paths: []string{"body"},
		}
	}

	result := make([]*pb.Message, 0, len(messages))
	for _, msg := range messages {
		rmsg := &pb.Message{}

		for _, path := range mask.Paths {
			switch path {
			case "body":
				rmsg.Body = msg.Body
			case "created_time":
				rmsg.CreatedTime = msg.CreatedTime
			case "title":
				rmsg.Title = msg.Title
			case "mentions":
				rmsg.Mentions = msg.Mentions
			default:
				return nil, status.Errorf(codes.InvalidArgument, "invalid field: %v", path)
			}
		}

		result = append(result, rmsg)
	}

	return result, nil
}
