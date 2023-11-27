package handlers

import (
	"net/http"

	"github.com/elue-dev/bookVerse/controllers"
	"github.com/elue-dev/bookVerse/helpers"
	"github.com/gorilla/mux"
)

func GetSingleUser (w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["id"]

	currUser, err := controllers.GetUser(userId)

	if err != nil {
		helpers.SendErrorResponse(w, 404, "User not found", err)
	}

	helpers.SendSuccessResponseWithData(w, 200, currUser)
}