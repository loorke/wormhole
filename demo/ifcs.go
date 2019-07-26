package main

import (
	"context"
	"time"
)

type Model struct {
	ID   string
	Time time.Time
}

type Greeter interface {
	Hello(ctx context.Context, q GreeterHelloReq) (GreeterHelloResp, error)
	Hello12(ctx context.Context, q GreeterHelloReq) (GreeterHelloResp, error)
}

type Greeter2 interface {
	Hello21(ctx context.Context, q GreeterHelloReq) (GreeterHelloResp, error)
	Hello22(ctx context.Context, q GreeterHelloReq) (GreeterHelloResp, error)
}

type GreeterHelloReq struct {
	Name        string
	NameChanged func(ctx context.Context, data string) (string, error)
}

type GreeterHelloResp struct {
	Name  string
	// Reply func(ctx context.Context, data string) (string, error) // TODO
}
