package handlers

import (
	"fmt"
	"net/http"

	"github.com/MostafaSensei106/GoWebBase/constants"
)

// HelloHandler handles the hello request.
// It first checks if the requested URL path is "/hello".
// If the path is not "/hello", it returns a 404 error to the client.
// It then checks if the request method is "GET".
// If the method is not "GET", it returns a "method is not supported" error to the client.
// If the method is "GET", it writes "Hello from GoWebBase" to the client.
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
