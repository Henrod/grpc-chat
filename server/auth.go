package main

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc"
)

func getAuth(token string) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Error(codes.Unauthenticated, "no auth sent")
		}

		t := md.Get("token")
		if len(t) == 0 {
			return nil, status.Error(codes.Unauthenticated, "no auth sent")
		}
		if t[0] != token {
			return nil, status.Error(codes.PermissionDenied, "failed to authenticate")
		}

		return handler(ctx, req)
	}
}

type auth struct {
	token string
}

func (a *auth) GetRequestMetadata(_ context.Context, _ ...string) (map[string]string, error) {
	return map[string]string{"token": a.token}, nil
}

func (a *auth) RequireTransportSecurity() bool { return true }
