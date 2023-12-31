// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package user

import (
	"context"

	"go-zero-microservice-demo/service/core/service"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AuthRequest      = service.AuthRequest
	AuthResponse     = service.AuthResponse
	CommonResponse   = service.CommonResponse
	GetTokenRequest  = service.GetTokenRequest
	GetTokenResponse = service.GetTokenResponse

	User interface {
		GetToken(ctx context.Context, in *GetTokenRequest, opts ...grpc.CallOption) (*GetTokenResponse, error)
		JwtAuth(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthResponse, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

func (m *defaultUser) GetToken(ctx context.Context, in *GetTokenRequest, opts ...grpc.CallOption) (*GetTokenResponse, error) {
	client := service.NewUserClient(m.cli.Conn())
	return client.GetToken(ctx, in, opts...)
}

func (m *defaultUser) JwtAuth(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	client := service.NewUserClient(m.cli.Conn())
	return client.JwtAuth(ctx, in, opts...)
}
