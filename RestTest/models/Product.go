package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ProductId	        string         `gorm:"type:varchar(64);unique;primaryKey" json:"product_id" form:"product_id"`
	ProductName 		string         `gorm:"type:varchar(80)" json:"product_name" form:"product_name"`
	BasicPrice			uint           `gorm:"type:varchar(100)" json:"basic_price" form:"basic_price"`
	CreatedDate   		time.Time      `json:"created_date" form:"created_date"`
}
