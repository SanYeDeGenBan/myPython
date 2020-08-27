package controller

import (
	"github.com/gin-gonic/gin"
	"go-order/entity"
)

func PlaceOrder(ctx *gin.Context) {
	var Order entity.Order
	err := ctx.ShouldBindJSON(&Order)
	if err != nil {
		panic("参数有误")
	}

}
