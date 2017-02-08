package other

import (
	"io"
	"net/http"
)

func checkPayment(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "dol checkPayment")
}

func incomePayment(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "dol incomePayment")
}
