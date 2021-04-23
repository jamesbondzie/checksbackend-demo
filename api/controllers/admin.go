package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"checksbackend/api/auth"
	"checksbackend/api/models"
	"checksbackend/api/responses"
)

//CreateNewLocationGroup is a func that will create a group
//by admin
func (s *Server) CreateNewLocationGroup(w http.ResponseWriter, r *http.Request) {
		
	 body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	location := models.UserLocation{}

	err = json.Unmarshal(body, &location)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	location.BeforeSave()
	
	err = location.ValidateData()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	userId, err := auth.ExtractTokenID(r)
	log.Println("[UID FROM ADMIN.GO]",userId)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}


	if userId != location.AdminID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}

	newLocationCreated, err := location.SaveAminSelectedLocation(s.DB)
	if err != nil {
		formattedError := responses.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, newLocationCreated.LocationID))
	responses.JSON(w, http.StatusCreated, newLocationCreated)
}