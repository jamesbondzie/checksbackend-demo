package auth

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)


//CreateToken is used for
//authentication and authorization
func CreateToken(userid uint32) (string, error) {
	
	var err error
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userid               //ref for user id
	atClaims["exp"] = time.Now().Add(time.Minute * 1).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil{
	   return "", err
	}
	return token, nil
  }

  

//PassportToken to be used as login
///sign pass by mobile app
//randon token
//func RandomKey()([]byte, error){


//	key := make([]byte,64)
//	key, err := getSecretKey()
//	if err != nil{
//		panic(err)
//	}
	
//	if _,err := rand.Read(key); err != nil{
//		panic(err)
//	}

//	block ,err := aes.NewCipher(key)
///	if err != nil{
//		return nil, err
//	}

//	aesgcm, err := cipher.NewGCM(block)
//	if err != nil{
//		return nil, err
//	}

//	iv := make([]byte, aesgcm.NonceSize())
//	if _, err := rand.Read(iv); err != nil{
//		return nil, err
//	}

//	key_ciphertext := aesgcm.Seal(iv, iv, key, nil)

//	fmt.Println("[ciphertext]", key_ciphertext)
	
//	fmt.Println("[KEY]", key)


//	return key, nil
//}



func Encrypt(key []byte)([]byte, error){

	secret_key, err := getSecretKey()
	if err != nil {
		panic(err)
	}
	
	block ,err := aes.NewCipher(secret_key)
	if err != nil{
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil{
		return nil, err
	}

	iv := make([]byte, aesgcm.NonceSize())
	if _, err := rand.Read(iv); err != nil{
		return nil, err
	}

	ciphertext := aesgcm.Seal(iv, iv, key, nil)

	//fmt.Println("[ciphertext]", ciphertext)

	return ciphertext, nil
}


func getSecretKey()([]byte, error){
	secret := os.Getenv("MASTER_KEY")

	if secret == " "{
		panic("Cannot get secret from .env")
	}

	secretBit, err := hex.DecodeString(secret)
	if err != nil{
		panic(err)
	}

	return secretBit, nil
}
 

//TokenValid func
func TokenValid(r *http.Request) error {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		Pretty(claims)
	}
	
	return nil
}

//ExtractToken from r
func ExtractToken(r *http.Request) string {
	keys := r.URL.Query()
	token := keys.Get("token")

	if token != "" {
		
		return token	
	}
	
	Bearer := r.Header.Get("Authorization")
	if len(strings.Split(Bearer, " ")) == 2 {
		return strings.Split(Bearer, " ")[1]
	}
	return " "
	
}

//ExtractTokenID after token extracter from header
func ExtractTokenID(r *http.Request) (uint32, error) {

	tokenString := ExtractToken(r)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		panic(err)
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id"]), 10, 32)
		if err != nil {
			return 0, err
		}
		return uint32(uid), nil
	}
	return 0, nil
	
}


//ExtractTokenReferenceSecret after token extracter from header
//func ExtractTokenReferenceSecret(r *http.Request) ([]byte, error) {

//	tokenString := ExtractToken(r)

//	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
//		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
//		}
//		return []byte(os.Getenv("ACCESS_SECRET")), nil
//	})
//	log.Println("token", token)
//	if err != nil {
//		panic(err)
//	}
//	claims, ok := token.Claims.(jwt.MapClaims)
//	if ok && token.Valid {
//		ref := fmt.Sprintf("%.0f", claims["ref"])
//		if err != nil {
//			return []byte{}, err
//		}
//		log.Println(ref)
//		return []byte(os.Getenv("ACCESS_SECRET")), nil
//	}
//	return  []byte{}, nil
//}


//Pretty display the claims licely in the terminal
func Pretty(data interface{}) {
	b, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(string(b))
}


