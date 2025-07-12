package logic

import (
	"context"
	"dcproject/dcrpc/basic/global"
	"dcproject/dcrpc/basic/models"
	"dcproject/dcrpc/dcrpc"
	"dcproject/dcrpc/internal/svc"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type CalcDistanceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCalcDistanceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CalcDistanceLogic {
	return &CalcDistanceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CalcDistanceLogic) CalcDistance(in *dcrpc.CalcDistanceRequest) (*dcrpc.CalcDistanceResponse, error) {
	// todo: add your logic here and delete this line
	// 默认单位
	unit := in.Unit
	if unit == "" {
		unit = "km"
	}

	dist, err := global.Client.GeoDist(l.ctx, "drivers", in.FromName, in.ToName, unit).Result()
	if err != nil {
		l.Logger.Errorf("计算 %s 到 %s 的距离失败: %v", in.FromName, in.ToName, err)
		return nil, err
	}

	//历程记录表
	orderdetail := models.LxhOrderDetails{
		OrderCode: in.OrderCode,
		TripKey:   in.TripKey,
		DriverKey: in.DriverKey,
	}
	err = global.DB.Create(&orderdetail).Error
	if err != nil {
		fmt.Println("订单详情成功")
	}

	//订单基本信息创建
	order := models.LxhOrders{
		OrderCode:   in.OrderCode,
		Amount:      float64(in.Amount),
		OrderStatus: in.OrderStatus,
		PassengerId: in.PassengerId,
		StartAddr:   in.FromName,
		EndEnd:      in.ToName,
		DriverId:    in.DriverId,
	}

	err = global.DB.Create(&order).Error

	if err != nil {
		fmt.Println("订单创建成功")
	}
	return &dcrpc.CalcDistanceResponse{
		Distance: dist,
	}, nil
}
