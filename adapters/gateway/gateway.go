package gateway

import (
	"fmt"
	"net/http"
)

var (
	adapters = make(map[string]Adapter)
)

type Adapter interface {
	GetId() string
	Pay(payer string, order string) string
}

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
