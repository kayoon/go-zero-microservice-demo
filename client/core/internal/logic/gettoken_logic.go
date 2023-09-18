package logic

import (
	"context"

	"go-zero-microservice-demo/client/core/internal/svc"
	"go-zero-microservice-demo/client/core/internal/types"
	"go-zero-microservice-demo/service/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GettokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGettokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GettokenLogic {
	return &GettokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GettokenLogic) Gettoken(req *types.GetTokenRequest) (resp *types.GetTokenResponse, err error) {
	uc, err := l.svcCtx.UserRpc.GetToken(l.ctx, &user.GetTokenRequest{
		UserId:   int32(req.UserId),
		GroupId:  req.GroupId,
		UserName: req.UserName,
	})
	if err != nil {
		return nil, err
	}
	resp = &types.GetTokenResponse{}
	resp.Token = uc.Token
	return
}
