package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MostafaSensei106/GoWebBase/server"
)

func main() {
	server.RegisterRoutes()

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
