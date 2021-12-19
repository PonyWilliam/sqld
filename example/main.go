package main

import (
	"fmt"
	"log"

	"github.com/PonyWilliam/sqld"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jmoiron/sqlx"
)

var Db *sqld.MySQL_D

type Person struct {
	UserId   int64  `db:"user_id"`
	Username string `db:"username"`
}

func init() {
	Db = sqld.SQL_init()
	Db.SetType("mysql")
	err, _ := Db.Connect("sh-cynosdbmysql-grp-123.sql.tencentcdb.com", 21270, "root", "root", "tests")
	if err != nil {
		fmt.Println("can not open mysql")
		log.Fatal(err)
	}
}

func main() {
	keys := []string{"username"}
	vals := []string{"William"}
	err, _ := Db.Insert("person", keys, vals)
	if err != nil {
		fmt.Println("无法插入数据")
		log.Fatal(err)
	}
	var person []Person
	err = Db.Select(&person, "select * from person")
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(person)
}
