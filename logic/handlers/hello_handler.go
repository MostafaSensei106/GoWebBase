package handlers

import (
	"fmt"
	"net/http"

	"github.com/MostafaSensei106/GoWebBase/constants"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != constants.HelloRoute {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello from GoWebBase")

}
