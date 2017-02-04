package main

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Money struct {
	CurrencyCode int `gorm:"column:currency_code"`
	Units        int `gorm:"column:unit"`
	Nanos        int `gorm:"column:nano"`
}

type CollectionOrder struct {
	OrderId     uint `sql:"AUTO_INCREMENT" gorm:"primary_key;column:order_id"`
	MerchantId  uint `gorm:"column:merchant_id"`
	CreatedDate time.Time `gorm:"column:create_date"`
	Money
}

func (o *CollectionOrder) TableName() string {
	return "collection_order"
}

func main() {
	db, err := gorm.Open("mysql", "root@tcp(10.0.3.87:3306)/billing?charset=utf8&parseTime=True")
	if err != nil {
		fmt.Errorf("Error while connectiong to db: %v", err)
	}
	defer db.Close()

	order := CollectionOrder{
		OrderId: 400000016,
	}
	db.Find(&order)
	fmt.Println(order)
}
