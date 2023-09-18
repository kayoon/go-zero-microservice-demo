package logic

import (
	"context"
	"go-zero-microservice-demo/service/user"

	"go-zero-microservice-demo/service/core/service"
	"go-zero-microservice-demo/service/internal/svc"
	"go-zero-microservice-demo/utils/common"

	"github.com/zeromicro/go-zero/core/logx"
)

type JwtAuthLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewJwtAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JwtAuthLogic {
	return &JwtAuthLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *JwtAuthLogic) JwtAuth(in *service.AuthRequest) (*service.AuthResponse, error) {
	uc, err := common.AnalyzeToken(in.Token)
	if err != nil {
		return nil, err
	}

	return &user.AuthResponse{
		UserId:     int32(uc.Id),
		GroupId:    uc.GroupId,
		UserName:   uc.Name,
		ExpireTime: "2023-01-01 00:00:00",
	}, nil
}
