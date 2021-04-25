package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"checksbackend/api/models"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //connect postgres DB
)

//Server struct for DB and Router
type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}


func (s *Server)CreateDB(DbUser, DbPassword, DbPort, DbHost,DbName string) {

	var err error
	DBURL := fmt.Sprintf("host=%s port=%s user=%s  sslmode=disable password=%s", DbHost, DbPort, DbUser,DbPassword)
	s.DB, err = gorm.Open("postgres", DBURL)
	if err != nil {
		panic(err)
	}
	defer s.DB.Close()
 
	s.DB = s.DB.Exec("CREATE DATABASE "+ DbName)
	if s.DB != nil {
		
		fmt.Println("Db created ", DbName)
	} else {
		fmt.Printf("Cannot create %s database",err)
	}

	s.DB = s.DB.Exec("USE "+ DbName)
	if s.DB  != nil {
		fmt.Printf("currently using %s database",DbName)
	}else{
		fmt.Printf("cannot use %s database",err)
	}

 }

 
//Initialize function to connect to postgres DB
//s is server
func (s *Server) InitializeDB(DbUser, DbPassword, DbPort, DbHost, DbName string) {

	var err error

	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	s.DB, err = gorm.Open("postgres", DBURL)
	if err != nil {
		fmt.Printf("Cannot connect to %s database",err)
		//log.Fatal("This is the error:", err)
	} else {
		fmt.Println("We are connected to the database", DbName)
	}

	s.DB.Debug().AutoMigrate(&models.User{}, &models.UserLocation{}) //database migration

	s.Router = mux.NewRouter()

	s.initializeRoutes()

}

//RunBackendResources function calls address and 
//serves
func (s *Server) RunBackendResources(addr string) {
	
	
	var ip = os.Getenv("ADDR")
	if ip == "" {
		log.Fatal("$PORT must be set")
	}
	 port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
		//port == "8080"
	}
	
	if port == " " && ip == "" {

		var ip = "localhost"
		  port = "8080"
		
		var  addr string  = ip + ":" + port
		fmt.Printf("Listening on %s",addr)

		//fmt.Println("Listening to port"+ " " + port)

	
}
	
	
//var  addr  = ip + ":" + port
log.Fatal(http.ListenAndServe(addr, s.Router))

}