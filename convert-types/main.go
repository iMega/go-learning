package main

import (
	"fmt"
	"strconv"
	"reflect"
)

type Money struct {
	// The 3-letter currency code defined in ISO 4217.
	CurrencyCode string `protobuf:"bytes,1,opt,name=currency_code,json=currencyCode" json:"currency_code,omitempty"`
	// The whole units of the amount.
	// For example if `currencyCode` is `"USD"`, then 1 unit is one US dollar.
	Units int64 `protobuf:"varint,2,opt,name=units" json:"units,omitempty"`
	// Number of nano (10^-9) units of the amount.
	// The value must be between -999,999,999 and +999,999,999 inclusive.
	// If `units` is positive, `nanos` must be positive or zero.
	// If `units` is zero, `nanos` can be positive, zero, or negative.
	// If `units` is negative, `nanos` must be negative or zero.
	// For example $-1.75 is represented as `units`=-1 and `nanos`=-750,000,000.
	Nanos int32 `protobuf:"varint,3,opt,name=nanos" json:"nanos,omitempty"`
}

type Order struct {
	Units uint
	Nanos uint
}

func main() {
	money := Money{
		CurrencyCode: "RUB",
		Units: 100,
		Nanos: 12,
	}

	order := Order{
		Units: uint(money.Units),
		Nanos: uint(money.Nanos),
	}

	fmt.Println(order)
	b, _ := strconv.ParseBool("true")
	d := reflect.ValueOf(b)
	c := strconv.FormatBool(true)
	z := reflect.ValueOf(c)
	fmt.Println(d.Kind(), z.Kind())
}