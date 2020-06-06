package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"google.golang.org/grpc/credentials"

	pb "github.com/Henrod/chat-example-2/protogen"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

const (
	certFile = "../creds/domain.crt"
	keyFile  = "../creds/domain.key"
)

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

	creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer(grpc.UnaryInterceptor(Metrics), grpc.Creds(creds))
	pb.RegisterFeedAPIServer(s, NewFeedAPI())

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
