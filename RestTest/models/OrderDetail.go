package models

import (
	"time"

	"gorm.io/gorm"
)

type OrderDetail struct {
	OrderDetailId      	string         `gorm:"primaryKey" json:"courier_id" form:"courier_id"`
	OrderId 			string         `gorm:"type:varchar(100);unique;not null" json:"company_name" form:"company_name"`
	ProductId			string         `gorm:"type:varchar(100);unique;not null" json:"company_name" form:"company_name"`
	PaymentMethodId		string		   `json:"created_at" form:"created_at"`
	CreatedDate			time.Time      `json:"created_at" form:"created_at"`
}
