package sqld

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func init() {
	//实例化部分功能
}

// type Mysql_D interface {
// 	Connect(host string, username string, password string, port int64, databse string) (error, bool)
// 	SetType(types string)
// 	Insert(table string) (error, bool)
// }

//type和isConnect作为私有变量不对外暴露
type MySQL_D struct {
	IsConnect bool
	Types     string
	Db        *sqlx.DB //可以通过这个直接使用sqlx的封装，因此更加具有灵活性
}

func SQL_init() *MySQL_D {
	var db *MySQL_D
	db = &MySQL_D{}
	return db
}

func (sql *MySQL_D) SetType(types string) {
	sql.Types = types
}

//Connect使用参数化，更加方便，全局变量在对象中进行封装
func (sql *MySQL_D) Connect(host string, port int64, username string, password string, databse string) (error, bool) {
	formats := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, databse)
	database, err := sqlx.Open("mysql", formats)
	if err != nil {
		return err, false
	}
	sql.Db = database
	sql.IsConnect = true
	return nil, true
}

func (sql *MySQL_D) Insert(table string, key []string, val []string) (error, bool) {
	//判断数据库是否已连接
	if !sql.IsConnect {
		return errors.New("Is not connect"), false
	}
	//判断是键值对是否相同
	if len(key) != len(val) {
		return errors.New(fmt.Sprintf("KV not compare,len_key:%d,len_v:%d", len(key), len(val))), false
	}
	formats := fmt.Sprintf("insert into %s(", table)
	for _, i := range key {
		if i != key[len(key)-1] {
			formats += (i + ", ")
			continue
		}
		formats += i
	}
	formats += ")values("
	for _, i := range val {
		if i != val[len(val)-1] {
			formats += ("'" + i + "'" + ", ")
			continue
		}
		formats += ("'" + i + "'")
		formats += ")"
	}
	fmt.Println(formats)
	_, err := sql.Db.Exec(formats)
	if err != nil {
		return err, false
	}
	return nil, true
}

func (sql *MySQL_D) Select(dest interface{}, query string, args ...interface{}) error {
	return sql.Db.Select(dest, query, args...)
}

func LogErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//直接执行
func (sql MySQL_D) Exec(command string) (sql.Result, error) {
	return sql.Db.Exec(command)
}
