package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	

	"checksbackend/api/auth"
	"checksbackend/api/models"
	"checksbackend/api/responses"
)

//Login func which takes token from
//scanned qrcode and submits to backend
//for authentication
func (s *Server) LoginTest(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	//do different validation for token here
	user.BeforeSave()
	err = user.Validate("login")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	token, err := s.SignIn(user.Email)
	if err != nil {
		formattedError := responses.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, token)
}


//SignIn func to be claaed in Login func
func (s *Server) SignIn(email string) (string, error) {

	var err error

	user := models.User{}

	//check for token in redis
	err = s.DB.Debug().Model(models.User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		return "", err
	}
	
	return auth.CreateToken(user.ID)
}
