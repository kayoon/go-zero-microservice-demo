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

func UserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserInfoRequest
		if err := httpx.Parse(r, &req); err != nil {
			err = errorx.NewErrorData(1002, "error", "参数异常")
			response.Response(w, &types.EmptyResponse{}, err)
			return
		}

		l := logic.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo(&req)
		response.Response(w, resp, err)
	}
}
