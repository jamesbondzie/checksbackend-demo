package controllers

import (
	"checksbackend/api/responses"
	
	"net/http"
)


//Home function is the default page for this api
func (s *Server) Home(w http.ResponseWriter, r *http.Request) {
	
	responses.JSON(w, http.StatusOK, "Welcome To Checks API")

}
