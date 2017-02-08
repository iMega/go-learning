package ya

import (
	"net/http"
	"github.com/imega/go-learning/adapters/gateway"
)

// Идентификаторы способов платежей
var payMethods = []string{"ac", "pc", "gp"}

type adapter struct {
	ShopId    string
	PayMethod string
}

func (adapter) Pay(payer string, order string) string {
	return "ya-adapter"
}

func (adapter) Handler(req gateway.Request) adapter {
	return adapter{}
}

// регистрация адаптера
func Invoke(srv *http.ServeMux) gateway.Adapter {
	srv.HandleFunc("/payments/yandex/check", checkPayment)
	srv.HandleFunc("/payments/yandex/incoming", incomePayment)

	a := adapter{
		ShopId: "sdf",
	}

	return a
}

// Получить идентификатор адаптера
func (a adapter) GetId () string {
	return "ya-adapter"
}

// получить способы платежей
func (a adapter) GetMethods() gateway.PayMethods {
	methods := gateway.PayMethods{
		a.GetId(): payMethods,
	}

	return methods
}