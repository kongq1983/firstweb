package dao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

var Db *sqlx.DB

func InitDb() (err error) {
	fmt.Printf("fmt init db ---------------------- \n")
	log.Printf("log init db ---------------------- \n")
	//dsn := "admin:admin123@tcp(127.0.0.1:23306)/mytest?charset=utf8mb4&parseTime=True"
	dsn := "admin:admin123@tcp(mysql.kq.com:23306)/mytest?charset=utf8&parseTime=True"
	Db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	return
}
