package responses

import (
	"errors"
	"strings"
)

//FormatError function
func FormatError(err string) error {

//	if strings.Contains(err, "firstname") {
//		return errors.New("Nickname Already Taken")
//	}

	if strings.Contains(err, "email") {
		return errors.New("Please, Email Already Taken")
	}

	if strings.Contains(err, "title") {
		return errors.New("Title Already Taken")
	}
//	if strings.Contains(err, "hashedPassword") {
//		return errors.New("Incorrect Password")
//	}
	return errors.New("Incorrect Details")
}



func Message(status bool, message string) (map[string]interface{}) {
	return map[string]interface{}{"status": status, "message": message}
}
