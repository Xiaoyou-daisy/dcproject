package logic

import (
	"context"
	"dcproject/dcrpc/basic/global"
	"fmt"
	"github.com/redis/go-redis/v9"
	"strconv"

	"dcproject/dcrpc/dcrpc"
	"dcproject/dcrpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type NearbyDriversLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewNearbyDriversLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NearbyDriversLogic {
	return &NearbyDriversLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}
func (l *NearbyDriversLogic) NearbyDrivers(in *dcrpc.NearbyDriversRequest) (*dcrpc.NearbyDriversResponse, error) {
	unit := in.Unit
	if unit == "" {
		unit = "km"
	}

	// 校验经纬度范围
	if in.Longitude < -180 || in.Longitude > 180 {
		return nil, fmt.Errorf("invalid longitude: %v (must be in [-180, 180])", in.Longitude)
	}
	if in.Latitude < -90 || in.Latitude > 90 {
		return nil, fmt.Errorf("invalid latitude: %v (must be in [-90, 90])", in.Latitude)
	}

	// 使用 GeoRadius 查询附近司机
	results, err := global.Client.GeoRadius(l.ctx, "drivers",
		in.Longitude, in.Latitude,
		&redis.GeoRadiusQuery{
			Radius:    in.Radius,
			Unit:      unit,
			WithDist:  true,
			WithCoord: true,
			Count:     int(in.Count),
			Sort:      "ASC",
		}).Result()
	if err != nil {
		l.Logger.Errorf("GeoRadius 查询失败: %v", err)
		return nil, fmt.Errorf("failed to get nearby drivers: %v", err)
	}

	// 如果没有找到司机，返回空的司机列表
	if len(results) == 0 {
		return &dcrpc.NearbyDriversResponse{Drivers: nil}, nil
	}

	// 构建司机数据返回列表
	var drivers []*dcrpc.Driver
	for _, loc := range results {
		// 解析 driver_id，确保 Name 是司机的 ID
		driverID, err := strconv.ParseInt(loc.Name, 10, 64)
		if err != nil {
			continue // 如果解析失败，跳过这个司机
		}

		// 添加司机到结果列表
		drivers = append(drivers, &dcrpc.Driver{
			DriverId:  driverID,
			Latitude:  loc.Latitude,
			Longitude: loc.Longitude,
			Distance:  loc.Dist,
		})
	}

	return &dcrpc.NearbyDriversResponse{
		Drivers: drivers,
	}, nil
}
