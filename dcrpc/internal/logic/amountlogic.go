package logic

import (
	"context"
	"dcproject/dcrpc/basic/global"
	"dcproject/dcrpc/basic/models"
	"dcproject/dcrpc/dcrpc"
	"dcproject/dcrpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AmountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

const (
	baseFare  = 14.0 // 起步价
	perKmFare = 10.0 // 每公里收费
)

func NewAmountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AmountLogic {
	return &AmountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AmountLogic) Amount(in *dcrpc.AmountRequest) (*dcrpc.AmountResponse, error) {
	// todo: add your logic here and delete this line
	// 获取从请求中获取的距离（单位为km）
	distance := in.Distance // 假设距离已通过其他服务或者计算得出

	// 计算总费用

	if distance <= 1 {
		// 距离在1公里内，使用起步价
		in.Amount = baseFare
	} else {
		// 超过1公里，按每公里收费计算
		in.Amount = baseFare + (distance-1)*perKmFare
	}
	global.DB.Model(&models.LxhOrders{}).
		Where("id = ?", in.Id).
		Updates(map[string]interface{}{
			"driver_id": in.DriverId,
			"amount":    in.Amount,
		})
	return &dcrpc.AmountResponse{
		Distance:    distance,
		TotalAmount: in.Amount,
	}, nil
}
