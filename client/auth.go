package main

import "context"

type auth struct {
	token string
}

func (a *auth) GetRequestMetadata(_ context.Context, _ ...string) (map[string]string, error) {
	return map[string]string{"token": a.token}, nil
}

func (a *auth) RequireTransportSecurity() bool { return true }
