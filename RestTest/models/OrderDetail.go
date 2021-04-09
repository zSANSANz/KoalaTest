package models

import (
	"time"
)

type OrderDetail struct {
	OrderDetailId      	string         `gorm:"type:varchar(64);primaryKey" json:"order_detail_id" form:"order_detail_id"`
	OrderId 			string         `gorm:"type:varchar(64);unique;not null" json:"company_name" form:"company_name"`
	ProductId			string         `gorm:"type:varchar(64);unique;not null" json:"company_name" form:"company_name"`
	Qty					uint		   `json:"created_at" form:"created_at"`
	CreatedDate			time.Time      `json:"created_at" form:"created_at"`
	Product   			Product	   	   `json:"product" form:"product"`
}
