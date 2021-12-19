package main

import (
	"fmt"
	"log"

	"github.com/PonyWilliam/sqld"
)

var Db sqld.MySQL_D

type Person struct {
	UserId   int64  `db:"user_id"`
	Username string `db:"username"`
}

func init() {
	Db.SetType("mysql")
	Db.Connect("sh-cynosdbmysql-grp-2o1mkprk.sql.tencentcdb.com", 21270, "123", "123", "/tests")
}

func main() {
	var person []Person
	err := Db.Select(&person, "select * from person")
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(person)
}
