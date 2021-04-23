package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"checksbackend/api/models"
	"checksbackend/api/responses"
)

//RegisterNewUser function to create/register new user
//for our service
//s -> server
func (s *Server) RegisterNewUser(w http.ResponseWriter, r *http.Request) {

	user := models.User{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}


	user.BeforeSave()


	err = user.Validate("register")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	userCreated, err := user.SaveUser(s.DB)
		
	if err != nil {	
		formattedError := responses.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	
	
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, userCreated.ID))
	log.Println("userCreated", userCreated)
	responses.JSON(w, http.StatusCreated, userCreated)

	//redirect User to home at frontEnd
	
}


