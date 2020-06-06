package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
)

func metrics() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		start := time.Now()

		resp, err = handler(ctx, req)
		log.Printf("method:%s, response_time:%v, error=%t", info.FullMethod, time.Since(start), err != nil)

		return resp, err
	}
}
