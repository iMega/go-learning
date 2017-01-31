package ya

import (
	"github.com/imega/go-learning/adapters/gateway"
	"net/http"
)

type YaAdapter struct {
	ShopId    string
	PayMethod string
}

func (YaAdapter) Pay(payer string, order string) string {
	return "ya-adapter"
}

func (YaAdapter) Handler(req gateway.Request) YaAdapter {
	return YaAdapter{}
}

func Invoke(srv *http.ServeMux) gateway.Adapter {
	srv.HandleFunc("/secure/payments/yandex/check_payment", checkPayment)
	srv.HandleFunc("/secure/payments/yandex/incoming_payment", incomePayment)

	a := YaAdapter{
		ShopId: "sdf",
	}
	return a
}

func (a YaAdapter) GetId () string {
	return "ya-adapter"
}