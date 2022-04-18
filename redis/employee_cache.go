package redis

import (
	"com.kq/dto"
	"encoding/json"
	"strconv"
	"time"

	//"encoding/json"
	"fmt"
)

func Get(id int) dto.Employee {

	var e dto.Employee

	var val, err = rdb.Get(ctx, "key").Result()
	if err != nil {
		//panic(err) //panic 让golang程序马上停止执行，慎用
	}
	fmt.Println("key", val)

	return e

}

func Set(e dto.Employee) {

	var json, err = json.Marshal(e)
	if err != nil {
		fmt.Printf("err=%v \n", err)
	}

	rdb.Set(ctx, intToString(e.Id), string(json), time.Hour*120)

}

func intToString(a int) string {
	return strconv.Itoa(a)
}
