package logic

import (
	"context"
	"dcproject/dcrpc/basic/global"
	"fmt"

	"dcproject/dcrpc/dcrpc"
	"dcproject/dcrpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLocalhostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLocalhostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLocalhostLogic {
	return &GetLocalhostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLocalhostLogic) GetLocalhost(in *dcrpc.GetLocalhostRequest) (*dcrpc.GetLocalhostResponse, error) {
	// 动态地根据请求中的 Name 查询
	result, err := global.Client.GeoPos(l.ctx, "drivers", in.Name).Result()
	if err != nil {
		l.Logger.Errorf("获取 %s 的地理位置失败: %v", in.Name, err)
		return nil, err
	}

	// 如果没有查询到
	if len(result) == 0 || result[0] == nil {
		return nil, fmt.Errorf("location %s not found", in.Name)
	}

	// 取第一个结果
	pos := result[0]
	l.Logger.Infof("Location: %s, Latitude: %f, Longitude: %f", in.Name, pos.Latitude, pos.Longitude)

	// 构造响应对象
	return &dcrpc.GetLocalhostResponse{
		Latitude:  pos.Latitude,
		Longitude: pos.Longitude,
		Name:      in.Name,
	}, nil
}
