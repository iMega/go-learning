package main

import (
	"fmt"
	"github.com/shopspring/decimal"
)

func main() {
	dec := decimal.New(1000, -2)
	price, _ := decimal.NewFromString("136.02")

	fmt.Println(dec.String())
	fmt.Println(price.Exponent())
}
