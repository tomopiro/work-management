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
	r.GET("/users", getUsers)

	r.Run(":80")
}

func gormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "quuta"
	PASS := "quuta"
	// "tcp('コンテナ名':'ポート番号'"
	PROTOCOL := "tcp(db:3306)"
	DBNAME := "quuta"

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

func getUsers(c *gin.Context) {
	db := gormConnect()
	defer db.Close()

	user := model.User{}
	db.Where("id = ?", "1").First(&user)
	c.String(http.StatusOK, user.Name+"\n")
}
