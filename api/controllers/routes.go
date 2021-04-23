package controllers

import "checksbackend/api/middlewares"



func(s *Server)initializeRoutes(){

	// Home Route
	s.Router.HandleFunc("/home", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")



	//WebsocketHandlers
	s.Router.HandleFunc("/ws-web", middlewares.SetMiddlewareJSON(s.WebAppSocketHandlerLogin)).Methods("GET")
	s.Router.HandleFunc("/ws-mobile", middlewares.SetMiddlewareJSON(wsHandlerMobileLogin)).Methods("GET")

	
	
	//SignUp route
	//s.Router.HandleFunc("/ws-registerNewUsersMobileClient", middlewares.SetMiddlewareJSON(s.RegisterNewUser)).Methods("POST")
	
	
	
	//Admin Create New Checkin Group route
	//s.Router.HandleFunc("/admin-create-new-group", middlewares.SetMiddlewareJSON(s.CreateNewLocationGroup)).Methods("POST")

	
	
	//http handler
	s.Router.HandleFunc("/mobile-login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	
	//s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.LoginTest)).Methods("POST")
}