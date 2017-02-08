package main

import (
	"net/http"

	"github.com/imega/go-learning/adapters/adapter/other"
	"github.com/imega/go-learning/adapters/adapter/ya"
	"github.com/imega/go-learning/adapters/gateway"
)

func main() {
	/*request := gateway.Request{
		Data: "",
	}*/

	/*gateway.Handle("ya-adapter", adapter.YaAdapter{
		ShopId: "12345",
	})*/

	//req := "qweqwe"
	//a :=

	//gateway.Handle("ya-adapter", request)
	//response := gateway.Listen("dol-adapter")

	mux := http.NewServeMux()
	gateway.Add(other.Invoke(mux))
	gateway.Add(ya.Invoke(mux))

	//account := Account{
	//	Id: "4004545",
	//}

	//account.AllowPayMethods(gateway.GetMethods())

	//mux.HandleFunc("/", gateway.AcceptPaymentHandler)

	//dol.Handler(mux)

	http.ListenAndServe("0.0.0.0:8080", mux)

	//fmt.Println(response)
}
