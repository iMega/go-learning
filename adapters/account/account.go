package main

import "github.com/imega/go-learning/adapters/gateway"

type Account struct {
	Id         string
	PayMethods gateway.PayMethods
}

func (a Account) AllowPayMethods(methods gateway.PayMethods) {

}
