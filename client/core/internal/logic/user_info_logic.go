package logic

import (
	"context"
	"go-zero-microservice-demo/service/user"

	"go-zero-microservice-demo/client/core/internal/svc"
	"go-zero-microservice-demo/client/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoRequest) (resp *types.UserInfoResponse, err error) {
	uc, err := l.svcCtx.UserRpc.JwtAuth(l.ctx, &user.AuthRequest{
		Token: req.Token,
	})
	if err != nil {
		return nil, err
	}
	resp = &types.UserInfoResponse{}
	resp.UserName = uc.UserName
	resp.GroupId = uc.GroupId
	return
}
