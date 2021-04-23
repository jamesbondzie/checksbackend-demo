package wesockets

import (
	"log"
	"net/http"


	"github.com/gorilla/websocket"
)





var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}



func WsUpgrade(w http.ResponseWriter, r *http.Request)(*websocket.Conn, error){

	
	//ws == c
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		m := "Unable to upgrade to websockets"
		log.Println(m)
		http.Error(w, m, http.StatusBadRequest)
		return c, err
	}

	if c != nil{
		m := "Client successfully connected..."
		log.Println("connected msg",m)
	}

	defer c.Close()
	
	return c, nil
}



func Writer(w http.ResponseWriter, r *http.Request)(conn *websocket.Conn){
	c, err := WsUpgrade(w,r)
	if err != nil {
		m := "Unable to upgrade to websockets in writer function"
		log.Println(m)
		http.Error(w, m, http.StatusBadRequest)
	}
	for{
		mt,m, err := c.ReadMessage()
		if err != nil{
			return
		}

		 err = c.WriteMessage(mt,m)
		 if err != nil {
			log.Println("writeERROR:", err)
			break
		}	
	}

	return c
}

