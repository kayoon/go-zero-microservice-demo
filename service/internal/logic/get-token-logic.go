package logic

import (
	"context"
	"go-zero-microservice-demo/utils/common"

	"go-zero-microservice-demo/service/core/service"
	"go-zero-microservice-demo/service/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTokenLogic {
	return &GetTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTokenLogic) GetToken(in *service.GetTokenRequest) (*service.GetTokenResponse, error) {
	token, err := common.GenerateToken(int(in.UserId), in.UserName, in.GroupId)
	if err != nil {
		return nil, err
	}

	return &service.GetTokenResponse{
		Token: token,
	}, nil
}
