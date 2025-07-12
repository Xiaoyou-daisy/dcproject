package global

import (
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

var DB *gorm.DB

var Client *redis.Client

var MoDB *mongo.Client

var MongoColl *mongo.Collection
