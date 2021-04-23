package api

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)


//LoadEnvVariables from .env file
func LoadEnvVariables() {

	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	} else {
		fmt.Println("Env values received successfully")
	}

}