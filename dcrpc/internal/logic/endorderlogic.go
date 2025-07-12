package logic

import (
	"context"

	"dcproject/dcrpc/dcrpc"
	"dcproject/dcrpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type EndOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEndOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EndOrderLogic {
	return &EndOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *EndOrderLogic) EndOrder(in *dcrpc.EndOrderRequest) (*dcrpc.EndOrderResponse, error) {
	// todo: add your logic here and delete this line

	return &dcrpc.EndOrderResponse{}, nil
}
