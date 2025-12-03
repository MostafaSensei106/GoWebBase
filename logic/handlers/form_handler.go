package handlers

import (
	"fmt"
	"net/http"

	"github.com/MostafaSensei106/GoWebBase/constants"
	"github.com/MostafaSensei106/GoWebBase/data/models"
)

// FormHandler handles the form submission request.
// It first checks if the form data is valid, by calling r.ParseForm().
// If the form data is invalid, it returns an error to the client.
// If the form data is valid, it creates a User model and prints the user's name and address to the client.
// It assumes that the form data contains the name and address fields.
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
