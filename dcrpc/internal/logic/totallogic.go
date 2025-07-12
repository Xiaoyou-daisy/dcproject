package logic

import (
	"context"
	"dcproject/dcrpc/basic/global"
	"dcproject/dcrpc/basic/models"
	"fmt"
	"go.uber.org/zap"
	"time"

	"dcproject/dcrpc/dcrpc"
	"dcproject/dcrpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type TotalLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTotalLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TotalLogic {
	return &TotalLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TotalLogic) Total(in *dcrpc.TotalRequest) (*dcrpc.TotalResponse, error) {
	// ----- 1. 解析并兼容多种格式 -----
	parse := func(str string) (time.Time, error) {
		if t, err := time.Parse("2006-01-02 15:04:05", str); err == nil {
			return t, nil
		}
		return time.Parse("2006-01-02", str)
	}

	start, err := parse(in.CreatedAt)
	if err != nil {
		zap.L().Error("开始时间解析失败", zap.String("value", in.CreatedAt), zap.Error(err))
		return nil, fmt.Errorf("开始时间格式错误，应为 YYYY-MM-DD 或 YYYY-MM-DD HH:MM:SS")
	}
	end, err := parse(in.EndAt)
	if err != nil {
		zap.L().Error("结束时间解析失败", zap.String("value", in.EndAt), zap.Error(err))
		return nil, fmt.Errorf("结束时间格式错误，应为 YYYY-MM-DD 或 YYYY-MM-DD HH:MM:SS")
	}
	// 包含结束日 23:59:59
	end = end.AddDate(0, 0, 1).Add(-time.Second)

	var orders []models.LxhOrders
	if err := global.DB.Where("pay_status = ? AND driver_id = ? AND created_at BETWEEN ? AND ?", in.PayStatus, in.DriverId, start, end).
		Find(&orders).Error; err != nil {
		zap.L().Error("订单查询失败", zap.Error(err))
		return nil, fmt.Errorf("订单查询失败")
	}

	var total float32
	for _, o := range orders {
		total += float32(o.Amount)
	}

	return &dcrpc.TotalResponse{
		Amount: total,
	}, nil
}
