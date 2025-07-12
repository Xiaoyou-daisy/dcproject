package logic

import (
	"context"
	"dcproject/dcrpc/basic/global"
	"dcproject/dcrpc/dcrpc"
	"dcproject/dcrpc/internal/svc"
	"fmt"
	"github.com/redis/go-redis/v9"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetLocalhostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetLocalhostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetLocalhostLogic {
	return &SetLocalhostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetLocalhostLogic) SetLocalhost(in *dcrpc.SetLocalhostRequest) (*dcrpc.SetLocalhostResponse, error) {
	// todo: add your logic here and delete this line
	// todo: add your logic here and delete this line
	//定义中国重要城市和省份（中文名）
	locations := []struct {
		Area      string
		latitude  float64
		longitude float64
	}{
		{"北京", 39.9042, 116.4074},   // 北京
		{"上海", 31.2304, 121.4737},   // 上海
		{"广州", 23.1291, 113.2644},   // 广州
		{"深圳", 22.5431, 114.0579},   // 深圳
		{"成都", 30.5728, 104.0668},   // 成都
		{"杭州", 30.2741, 120.1551},   // 杭州
		{"武汉", 30.5928, 114.3055},   // 武汉
		{"重庆", 29.5630, 106.5516},   // 重庆
		{"天津", 39.3434, 117.3616},   // 天津
		{"南京", 32.0603, 118.7969},   // 南京
		{"惠南", 30.9791, 121.7600},   // 惠南
		{"惠南周围", 30.9790, 121.7601}, // 惠南周围
	}

	// 将这些位置存入Redis
	for i, loc := range locations {
		driverID := int64(i + 1) // 可以根据实际情况生成更合理的 driverID
		_, err := global.Client.GeoAdd(l.ctx, "drivers", &redis.GeoLocation{
			Name:      strconv.FormatInt(driverID, 10), // 将 driverID 转换为字符串存储
			Longitude: loc.longitude,
			Latitude:  loc.latitude,
		}).Result()
		if err != nil {
			l.Logger.Errorf("添加位置到Redis失败: %v", err)
			return nil, err
		}
	}

	return &dcrpc.SetLocalhostResponse{
		Msg: fmt.Sprintf("城市 %s 已成功添加。", in.Name),
	}, nil
}
