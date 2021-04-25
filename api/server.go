package api

import (
	"os"

	"checksbackend/api/controllers"
)



var s = controllers.Server{}

//init function to;
//Load .env.local variables
//Connect redis server
//Connect Postgres DB

func init() {
	
	LoadEnvVariables()
	s.CreateDB(os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	s.InitializeDB(os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	//RedisServer()


}


//Run function to call server
func Run(){
	
	
	
	s.RunBackendResources(":8080")
}