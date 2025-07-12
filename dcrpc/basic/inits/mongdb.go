package inits

import (
	"context"
	"dcproject/dcrpc/basic/global"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var err error

func MongeDB() {
	// 设置客户端选项
	clientOptions := options.Client().ApplyURI("mongodb://mongodb:2003225zyh@14.103.243.149:27017/?authSource=admin")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 连接到MongoDB
	global.MoDB, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	// 选择数据库和集合
	global.MoDB.Database("local").Collection("coll") // 替换为你的库/集合名

	// 选择数据库和集合（真正赋值）
	global.MongoColl = global.MoDB.Database("admin").Collection("数据源")

	// 创建地理位置索引（可选）
	index := mongo.IndexModel{
		Keys: bson.M{"location": "2dsphere"},
	}
	_, err = global.MongoColl.Indexes().CreateOne(ctx, index)
	if err != nil {
		log.Println(" 创建索引失败:", err)
	}

}
