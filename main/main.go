package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"go-order/dao"
	"net/http"
)

func init() {
	var err error
	dao.DataBaseConnector, err = gorm.Open("mysql", "")
	if err != nil {
		panic("数据库初始化失败")
	}
}

func main() {
	engine := gin.Default()
	engine.GET("/test", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"data": "Success"})
	})

	err := engine.Run("920429")
	if err != nil {
		panic("服务启动失败")
	}
}
