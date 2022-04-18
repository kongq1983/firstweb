package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
)

//var rdb = redis.NewClient(&redis.Options{
//	Addr:     "mysql.kq.com:26379",
//	Password: "admin123456", // no password set
//	DB:       0,  // use default DB
//})

var ctx = context.Background()

// 定义一个全局变量
var rdb *redis.Client

func InitRedis() (err error) {

	rdb = redis.NewClient(&redis.Options{
		Addr:     "mysql.kq.com:26379", // 指定
		Password: "admin123456",
		DB:       0, // redis一共16个库，指定其中一个库即可
	})
	_, err = rdb.Ping(ctx).Result()

	if err != nil {
		log.Printf("redis  is stop \n")
		panic(err)
	}

	fmt.Printf("redis is connnected ! ")

	return
}
