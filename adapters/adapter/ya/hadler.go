package ya

import (
	"net/http"
	"io"
)

func checkPayment(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "ya checkPayment")
}

func incomePayment(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "ya incomePayment")
}
