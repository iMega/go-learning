package gateway

import (
	"fmt"
	"net/http"
)

var (
	adapters   = make(map[string]Adapter)
	payMethods = make(map[string][]string)
)

type Adapter interface {
	GetId() string
	//Pay(payer string, order string) string
	//GetMethods() PayMethods
}

//type PayMethods map[string][]string

type Request struct {
	Data string
}

func Handle(a Adapter) {
	///adapters[AdapterId] = a
	fmt.Println(adapters)
}

func Pay() {

}

func AcceptPaymentHandler(w http.ResponseWriter, req *http.Request) {

}

func Listen(AdapterId string) string {
	adapter := adapters[AdapterId]
	return adapter.Pay("pay", "order")
}

func Add(a Adapter) {
	adapters[a.GetId()] = a
}

func GetMethods() PayMethods {
	for id, adapter := range adapters {
		m := adapter.GetMethods()
		payMethods[id] = m[id]
	}
	return payMethods
}
