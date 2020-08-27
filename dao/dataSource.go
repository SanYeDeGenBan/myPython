package dao

import (
	"github.com/jinzhu/gorm"
	"go-order/entity"
)

var DataBaseConnector *gorm.DB

func OrderInsert(order *entity.Order) {
	DataBaseConnector.NewRecord(order)
}
