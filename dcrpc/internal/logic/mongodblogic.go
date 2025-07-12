package logic

import (
	"context"
	"dcproject/dcrpc/basic/global"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"time"

	"dcproject/dcrpc/dcrpc"
	"dcproject/dcrpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type MongoDBLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMongoDBLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MongoDBLogic {
	return &MongoDBLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MongoDBLogic) MongoDB(in *dcrpc.MongoDBRequest) (*dcrpc.MongoDBResponse, error) {
	// todo: add your logic here and delete this line
	geoPos, err := global.Client.GeoPos(l.ctx, in.GeoKey, in.DriverIds...).Result()
	if err != nil {
		return nil, err
	}

	updated := 0
	for i, id := range in.DriverIds {
		pos := geoPos[i]
		if pos == nil {
			continue
		}
		doc := bson.M{
			"driverId": id,
			"location": bson.M{
				"type":        "Point",
				"coordinates": []float64{pos.Longitude, pos.Latitude},
			},
			"updatedAt": time.Now(),
		}
		// upsert
		_, err := l.svcCtx.MongoColl.UpdateOne(
			l.ctx,
			bson.M{"driverId": id},
			bson.M{"$set": doc},
			options.Update().SetUpsert(true),
		)
		if err != nil {
			zap.S().Error("失败")
			continue
		}
		updated++
	}

	return &dcrpc.MongoDBResponse{
		Updated: int64(updated),
	}, nil
}
