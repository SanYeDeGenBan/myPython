package entity

import "github.com/jinzhu/gorm"

type Order struct {
	gorm.Model
	OrderId string `json:"orderId"`
}
