package handlers

import (
	"fmt"
	"net/http"

	"github.com/MostafaSensei106/GoWebBase/constants"
	"github.com/MostafaSensei106/GoWebBase/data/models"
)

func FormHandler(w http.ResponseWriter, r *http.Request) {
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
