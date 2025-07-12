package pkg

import (
	"context"
	"github.com/redis/go-redis/v9"
)

const geoKey = "上海浦东新区"

type RedisGeo struct {
	client *redis.Client
}

func NewRedisGeo(client *redis.Client) *RedisGeo {
	return &RedisGeo{client: client}
}

// 添加或更新位置
func (g *RedisGeo) AddLocation(ctx context.Context, userID string, lat, lng float64) error {
	location := &redis.GeoLocation{
		Name:      userID,
		Longitude: lng,
		Latitude:  lat,
	}

	_, err := g.client.GeoAdd(ctx, geoKey, location).Result()
	return err
}

// 获取位置坐标
func (g *RedisGeo) GetLocation(ctx context.Context, userID string) (lat, lng float64, err error) {
	positions, err := g.client.GeoPos(ctx, geoKey, userID).Result()
	if err != nil {
		return 0, 0, err
	}
	if len(positions) == 0 || positions[0] == nil {
		return 0, 0, nil
	}

	return positions[0].Latitude, positions[0].Longitude, nil
}

// 查询附近位置
func (g *RedisGeo) Nearby(ctx context.Context, lat, lng, radius float64, unit string) ([]redis.GeoLocation, error) {
	query := &redis.GeoRadiusQuery{
		Radius:      radius,
		Unit:        unit,
		WithCoord:   true,
		WithDist:    true,
		WithGeoHash: false,
		Sort:        "ASC",
	}

	return g.client.GeoRadius(ctx, geoKey, lng, lat, query).Result()
}

// 计算两点距离
func (g *RedisGeo) Distance(ctx context.Context, userID1, userID2, unit string) (float64, error) {
	return g.client.GeoDist(ctx, geoKey, userID1, userID2, unit).Result()
}
