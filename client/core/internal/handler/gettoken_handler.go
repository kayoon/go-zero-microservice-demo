package handler

import (
	"go-zero-microservice-demo/utils/errorx"
	"net/http"

	"go-zero-microservice-demo/utils/response"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-microservice-demo/client/core/internal/logic"
	"go-zero-microservice-demo/client/core/internal/svc"
	"go-zero-microservice-demo/client/core/internal/types"
)

func GettokenHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetTokenRequest
		if err := httpx.Parse(r, &req); err != nil {
			err = errorx.NewErrorData(1002, "error", "参数异常")
			response.Response(w, &types.EmptyResponse{}, err)
			return
		}

		l := logic.NewGettokenLogic(r.Context(), svcCtx)
		resp, err := l.Gettoken(&req)
		response.Response(w, resp, err)
	}
}
