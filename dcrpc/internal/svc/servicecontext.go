package svc

import (
	"context"
	"dcproject/dcrpc/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ServiceContext struct {
	Config    config.Config
	MongoColl *mongo.Collection // 加这个字段
}

func NewServiceContext(c config.Config) *ServiceContext {

	// 初始化 Mongo 客户端
	clientOpts := options.Client().ApplyURI("mongodb://14.103.243.149:27017") // 替换为你的地址
	client, err := mongo.Connect(context.Background(), clientOpts)
	if err != nil {
		panic(err)
	}

	// 选择数据库和集合
	coll := client.Database("driver").Collection("database") // 替换为你的库/集合名

	return &ServiceContext{
		Config:    c,
		MongoColl: coll,
	}
}
