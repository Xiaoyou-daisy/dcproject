package inits

import (
	"dcproject/dcrpc/basic/global"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func ExampleClient() {
	global.Client = redis.NewClient(&redis.Options{
		Addr:     "14.103.243.149:6379",
		Password: "2003225zyh", // no password set
		DB:       0,            // use default DB
	})
	fmt.Println("redis init success")

}
