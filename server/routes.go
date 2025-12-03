package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MostafaSensei106/GoWebBase/constants"
	"github.com/MostafaSensei106/GoWebBase/logic/handlers"
)

func Execute() {
	// Static files
	fileServer := http.FileServer(http.Dir(constants.StaticFolder))
	http.Handle(constants.IndexRoute, fileServer)

	// Handlers
	http.HandleFunc(constants.FormRoute, handlers.FormHandler)
	http.HandleFunc(constants.HelloRoute, handlers.HelloHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
