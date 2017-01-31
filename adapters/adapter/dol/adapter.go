package dol

import (
	"net/http"
	"github.com/imega/go-learning/adapters/gateway"
)

type dolAdapter struct {
	Project  string
	ModeType string
}

func (a dolAdapter) Pay(payer string, order string) string {
	return "dol-adapter="
}

func Invoke(srv *http.ServeMux) gateway.Adapter {
	srv.HandleFunc("/secure/payments/dengionline/check_payment", checkPayment)
	srv.HandleFunc("/secure/payments/dengionline/incoming_payment", incomePayment)

	a := dolAdapter{
		Project: "sdf",
		ModeType: "fdgdg",
	}
	return a
}

func (a dolAdapter) GetId () string {
	return "dol-adapter"
}