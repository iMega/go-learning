package main

import (
	"github.com/imega/go-learning/adapters/gateway"
	//"fmt"
	"github.com/imega/go-learning/adapters/adapter/dol"
	"net/http"
	"github.com/imega/go-learning/adapters/adapter/ya"
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
	gateway.Add(dol.Invoke(mux))
	gateway.Add(ya.Invoke(mux))
	//mux.HandleFunc("/", gateway.AcceptPaymentHandler)

	//dol.Handler(mux)

	http.ListenAndServe("0.0.0.0:8080", mux)

	//fmt.Println(response)
}
