package main

import (
	"net/http"

	"server/model"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := gin.Default()

	r.GET("/hello", hello)
	r.GET("/employees", getEmployees)

	r.Run(":80")
}

func gormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "test"
	PASS := "pass"
	// "tcp('コンテナ名':'ポート番号'"
	PROTOCOL := "tcp(db:3306)"
	DBNAME := "work"

	// Datetime→time.Timeへの変換のため?parseTime=true を入れる必要あり
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
	db, err := gorm.Open(DBMS, CONNECT)
	db.LogMode(true)

	if err != nil {
		panic(err.Error())
	}
	return db
}

func hello(c *gin.Context) {
	c.String(http.StatusOK, "Hello world\n")
}

func getEmployees(c *gin.Context) {
	db := gormConnect()
	defer db.Close()

	employee := model.Employee{}
	db.Where("id = ?", "0001").First(&employee)
	c.String(http.StatusOK, employee.Name+"\n")
}
