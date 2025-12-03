package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MostafaSensei106/GoWebBase/constants"
	"github.com/MostafaSensei106/GoWebBase/data/models"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
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

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err %v", err)
		return
	}

	fmt.Fprintf(w, "POST Request Is Sucessful")
	user := models.User{
		Name:    r.FormValue(constants.UserFormName),
		Address: r.FormValue(constants.UserFormAddress),
	}
	fmt.Fprintf(w, "Name = %s\n", user.Name)
	fmt.Fprintf(w, "Address = %s\n", user.Address)

}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle(constants.IndexRoute, fileServer)
	http.HandleFunc(constants.FormRoute, formHandler)
	http.HandleFunc(constants.HelloRoute, helloHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
