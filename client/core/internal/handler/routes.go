// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"go-zero-microservice-demo/client/core/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/token/gettoken",
				Handler: GettokenHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/userinfo",
				Handler: UserInfoHandler(serverCtx),
			},
		},
	)
}
