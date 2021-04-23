package controllers

import (
	"time"

	"fmt"
	"io"
	"log"
	"net/http"
	"encoding/json"
	

	"github.com/gorilla/websocket"
)




type Message struct {
	ID 		string `json:"id"`
	Token 	string `json:"token"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	CurrentLocation time.Time `json:"current_location"`
}

//type Client struct{
//	ID int `json:"id"`
//	client  map[*websocket.Conn]bool
//}


var client = make(map[*websocket.Conn]bool)
var send = make(chan *Message)
var broadcast = make(chan *Message)
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}



func writer(token *Message) {
	send <- token 
}


func (s *Server) Login(w http.ResponseWriter, r *http.Request) {

	var pass Message

	
	if err := json.NewDecoder(r.Body).Decode(&pass); err != nil {
		if websocket.IsCloseError(err, websocket.CloseGoingAway) || err == io.EOF {
			log.Println("Websocket closed!")
			return
		}
	}

	//decrypt token here 

	defer r.Body.Close()
	go writer(&pass)
	 
}



func wsHandlerMobileLogin(w http.ResponseWriter, r *http.Request) {

	//ws == c
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		m := "Unable to upgrade to websockets"
		log.Println(m)
		http.Error(w, m, http.StatusBadRequest)
		return
	}

	if c != nil {
		m := "Client successfully connected...ready to handle login post request"
		log.Println("connected msg", m)
	}


	defer c.Close()

	for {
		val   := <-send
		//token := fmt.Sprintf(val.ID, val.Token, val.CreatedAt.Format(time.Now().Format("01-02-2006 15:04:05")))
		t := time.Now().Local().UTC().Format("02-01-2006 15:04:05")
		if err != nil {
			fmt.Println(err)
		}
		data := fmt.Sprint(val.Token, " " , t, "", "CHECKOUT")
						
		err = c.WriteJSON(data)
		if err != nil {
			log.Println("writeERROR:", err)
			return
		}	
	}
}

