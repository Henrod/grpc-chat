package main

import (
	"context"
	"log"
	"net"
	"net/http"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"

	"google.golang.org/grpc/credentials"

	pb "github.com/Henrod/chat-example-2/protogen"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

const (
	certFile = "../creds/domain.crt"
	keyFile  = "../creds/domain.key"
	token    = "token"
)

func startHTTPServer() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()

	creds, err := credentials.NewClientTLSFromFile(certFile, "")
	if err != nil {
		return err
	}

	opts := []grpc.DialOption{grpc.WithTransportCredentials(creds)}
	if err := pb.RegisterFeedAPIHandlerFromEndpoint(ctx, mux, "localhost:8000", opts); err != nil {
		return err
	}

	return http.ListenAndServeTLS("localhost:8001", certFile, keyFile, mux)
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

	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			metrics(),
			auth(token),
		)),
		grpc.Creds(creds),
	)
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
