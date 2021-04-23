package middlewares

import (
	"fmt"
	"net/http"
	"time"
)

//Message sent to us by the flutter web client
type Message struct {
	Latitude 	string `json:"lat"`
	Longitude   string `json:"long"`
}


type TokenData struct {
	Token      string `json:"token"`
}

//SetMiddlewareJSON header json
func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}



// ValidateMessage so that we know it's valid JSON 
//and contains a Longitude and Latitude
func ValidateMessage(data []byte) (Message, error) {
	var msg Message

	
	//const LATITUDE_PATTERN="^(\\+|-)?(?:90(?:(?:\\.0{1,6})?)|(?:[0-9]|[1-8][0-9])(?:(?:\\.[0-9]{1,6})?))$";
    //const LONGITUDE_PATTERN="^(\\+|-)?(?:180(?:(?:\\.0{1,6})?)|(?:[0-9]|[1-9][0-9]|1[0-7][0-9])(?:(?:\\.[0-9]{1,6})?))$";

	//if err := json.Unmarshal(data, &msg); err != nil {
	//	return msg, errors.Wrap(err, "Unmarshaling message")
	//}
	
	//if msg.Latitude != "" {
	//	return msg, errors.New("Latitude  can't be empty")
	//}
	//if msg.Longitude == "" {
	//	return msg, errors.New("Longitude cordinates can't be empty")
	//}

	//if msg.Latitude != "" && msg.Latitude != LATITUDE_PATTERN{
	//	return msg, errors.New("Latitude  not correct")
	//}
	//if msg.Longitude != "" && msg.Longitude != LONGITUDE_PATTERN{
	//	return msg, errors.New("Longtitude  not correct")
	//}

	//do more validation here

	
	return msg, nil
}

//func ValidateLoginToken(tokenData []byte)(TokenData, error){
//	 var tokenMsg TokenData
//	 var tokenMsgLen = 32
//		var err = "Token appears to be invalide"
//	 if  len(tokenData) != tokenMsgLen{
//		panic(err)
		//return
//	 }
//	
//	return tokenMsg, nil
//}

//DoEvery auto call createToken func every second
func DoEvery(d time.Duration, f func(time.Time)){

	for x := range time.Tick(d){
		f(x)
	}
}



func HelloWorld(t time.Time){
	fmt.Println("%v: Hello world\n ",t)
}



func dosomething( s string){
	fmt.Println("doing something 2" , s)
}


func StartPolling2(){
	for{

		<-time.After(2*time.Second)
		go dosomething("polling from 2")
	}
}