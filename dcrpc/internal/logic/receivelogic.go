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

type ReceiveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewReceiveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReceiveLogic {
	return &ReceiveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ReceiveLogic) Receive(in *dcrpc.ReceiveRequest) (*dcrpc.ReceiveResponse, error) {
	// todo: add your logic here and delete this line

	//确认司机行程和用户的行程一致

	//修改订单司机的Id
	var order models.LxhOrders
	global.DB.Model(&order).Where("id = ?", in.Id).Update("driver_id", in.DriverId)

	//创建审核
	var OrderPickup = models.OrderPickups{
		PickupId: in.PickupId,
		OrderId:  order.Id,
		DriverId: in.DriverId,
		Status:   in.Status,
	}
	err := global.DB.Create(&OrderPickup).Error

	if err != nil {
		return nil, fmt.Errorf("司机接取订单失败")
	}

	return &dcrpc.ReceiveResponse{Msg: "司机接取订单"}, nil
}
