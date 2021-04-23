package controllers

import (
	//"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"checksbackend/api/auth"
	"checksbackend/api/middlewares"
	//"checksbackend/api/models"

	"github.com/gorilla/websocket"
)






var Upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}


//WebAppSocketHandlerLogin function handlers 
//all websocket connection from the client
//;browser -checks webapp
func (s *Server) WebAppSocketHandlerLogin(w http.ResponseWriter, r *http.Request) {
	
	//dosomething()

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	c, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		m := "Unable to upgrade to websockets"
		log.Println(m)
		http.Error(w, m, http.StatusBadRequest)
		return
	}

	if c != nil{
		m := "Client successfully connected..."
		log.Println("connected msg",m)
	}
	
	defer c.Close()
	
	for{
		log.Println("msg from loop" )
		mt, data, err := c.ReadMessage()
		log.Println("data read", data)

		if err != nil {
			if websocket.IsCloseError(err, websocket.CloseGoingAway) || err == io.EOF {
				log.Println("Websocket closed!")
				break
			}
			log.Print("Error reading websocket message")
	    } 	
		switch mt{
			case websocket.TextMessage:
				  //do validattion
				  m,err := middlewares.ValidateMessage(data)
				  if err != nil {
					log.Println("Error from data validation", err)
					break
				}else{

					var msgJSON = m
					
					log.Println("msg from else state:", msgJSON)
					
					for{
						 key := make([]byte,32)
                         //var ID uint32 = 1
						<-time.After(3*time.Second)
						m, err := auth.Encrypt(key)

						if err != nil{
							log.Printf("msgJSON err: %s", err)
						}					
						err = c.WriteJSON(m)
						if err != nil {
							log.Println("writeERROR:", err)
							break
						}			
					}		
				
				}
				
			default:
				log.Println("Unknow message")
		}

	}

}


//func dosomething(){
//	log.Println("from dosomething")
//}


//WebSignIn func 
//check long, lat from DB 
//and generate token
//func (s *Server) WebSignIn(longitude float64) (string, error) {

//	var err error

//	location := models.UserLocation{}

	//check for token in redis
//	err = s.DB.Debug().Model(models.UserLocation{}).Where("longitude = ?", longitude).Take(&location).Error
//	if err != nil {
//		return "",err
//	}
	
	
//return auth.CreateToken(location.AdminID)

//}



  //func AutoCallPassport(){
	
//	var userID  = 1
//	for{
		
//		<-time.After(2*time.Second)
//		go auth.CreateToken(uint32(userID))
//	}
//}