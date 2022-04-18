package main

import (
	"com.kq/myweb"
	"com.kq/redis"
)

func main() {

	redis.InitRedis()

	myweb.Init()

}
