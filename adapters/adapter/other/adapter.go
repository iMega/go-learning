package other

import (
	"net/http"
	"github.com/imega/go-learning/adapters/gateway"
)

// Идентификаторы способов платежей
var payMethods = []string{"142"}

type otherAdapter struct {
	Project  string
	ModeType string
}

func (a otherAdapter) Pay(payer string, order string) string {
	return "other-adapter="
}

func Invoke(srv *http.ServeMux) gateway.Adapter {
	srv.HandleFunc("/payments/check", checkPayment)
	srv.HandleFunc("/payments/incoming", incomePayment)

	a := otherAdapter{
		Project:  "sdf",
		ModeType: "fdgdg",
	}
	return a
}

func (a otherAdapter) GetId() string {
	return "other-adapter"
}

func (a otherAdapter) GetMethods() gateway.PayMethods {
	methods := gateway.PayMethods{
		a.GetId(): payMethods,
	}

	return methods
}
