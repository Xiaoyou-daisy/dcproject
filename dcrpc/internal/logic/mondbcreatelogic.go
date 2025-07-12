package logic

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"

	"dcproject/dcrpc/dcrpc"
	"dcproject/dcrpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type MonDBCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMonDBCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MonDBCreateLogic {
	return &MonDBCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MonDBCreateLogic) MonDBCreate(in *dcrpc.MonDBCreateRequest) (*dcrpc.MonDBCreateResponse, error) {
	// todo: add your logic here and delete this line
	// 创建 2dsphere 索引（只执行一次）
	_, err := l.svcCtx.MongoColl.Indexes().CreateOne(l.ctx,
		mongo.IndexModel{
			Keys: bson.D{{Key: "location", Value: "2dsphere"}},
		},
	)
	if err != nil {
		l.Logger.Error("创建索引失败:", err)
		// 不是致命错误，可以继续插入
	}

	var failed []string
	successCount := 0

	for _, city := range in.Cities {
		doc := bson.M{
			"name": city.Name,
			"location": bson.M{
				"type":        "Point",
				"coordinates": []float64{city.Longitude, city.Latitude},
			},
			"createdAt": time.Now(),
		}

		_, err := l.svcCtx.MongoColl.InsertOne(l.ctx, doc)
		if err != nil {
			l.Logger.Errorf("插入城市 %s 失败: %v", city.Name, err)
			failed = append(failed, city.Name)
			continue
		}

		successCount++
	}

	return &dcrpc.MonDBCreateResponse{
		Inserted: int32(successCount),
		Failed:   failed,
	}, nil
}
