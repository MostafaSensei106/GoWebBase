package server

import (
	"net/http"

	"github.com/MostafaSensei106/GoWebBase/constants"
	"github.com/MostafaSensei106/GoWebBase/logic/handlers"
)

func RegisterRoutes() {
	// Static files
	fileServer := http.FileServer(http.Dir(constants.StaticFolder))
	http.Handle(constants.IndexRoute, fileServer)

	// Handlers
	http.HandleFunc(constants.FormRoute, handlers.FormHandler)
	http.HandleFunc(constants.HelloRoute, handlers.HelloHandler)
}
