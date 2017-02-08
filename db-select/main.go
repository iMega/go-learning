package main

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"crypto/rand"
	//"github.com/shopspring/decimal"
	//"strconv"
	//"golang.org/x/text/currency"
)

type Money struct {
	CurrencyCode int
	Units        int
	Nanos        int
}

type Order1 struct {
	OrderId     uint
	MerchantId  uint
	CreatedDate time.Time
	Money
}

func randToken() string {
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

type Order struct {
	OrderId      uint `sql:"AUTO_INCREMENT" gorm:"primary_key;column:order_id"`
	MerchantId   uint `gorm:"column:merchant_id"`
	CreatedDate  time.Time `gorm:"column:create_date"`
	CurrencyCode int `gorm:"column:currency_code"`
	Amount       string `gorm:"column:amount"`
}

func main() {
	db, err := gorm.Open("mysql", "root@tcp(10.0.3.87:3306)/billing?charset=utf8&parseTime=True")
	if err != nil {
		fmt.Errorf("Error while connectiong to db: %v", err)
	}
	defer db.Close()

	/*order := Order1{
		MerchantId:   12,
		CreatedDate:  time.Now(),
		Money: Money{
			CurrencyCode: 643,
			Units:        200,
			Nanos:        1234,
		},
	}
	dec, _ := decimal.NewFromString(fmt.Sprintf("%s.%s", strconv.Itoa(order.Units), strconv.Itoa(order.Nanos)))
*/
	/*orderTmp := struct {
		OrderId      uint `sql:"AUTO_INCREMENT" gorm:"primary_key;column:order_id"`
		MerchantId   uint `gorm:"column:merchant_id"`
		CreatedDate  time.Time `gorm:"column:create_date"`
		CurrencyCode int `gorm:"column:currency_code"`
		Amount       string `gorm:"column:amount"`
		gorm.JoinTableHandler
	}{
		MerchantId:   order.MerchantId,
		CreatedDate:  order.CreatedDate,
		CurrencyCode: order.CurrencyCode,
		Amount:       dec.String(),
		JoinTableHandler: gorm.JoinTableHandler{
			TableName: "orders",
		},
	}*/
/*
	orderTmp := Order{
		MerchantId:   order.MerchantId,
		CreatedDate:  order.CreatedDate,
		CurrencyCode: order.CurrencyCode,
		Amount:       dec.String(),
	}

currency.EUR
	trn := db.Begin()

	if err := trn.Create(&orderTmp).Error; err != nil {
		trn.Rollback()
		fmt.Println(err)
	}

	trn.Commit()*/



}
