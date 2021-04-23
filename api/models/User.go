package models

import (
	"errors"
	"html"
	"strings"
	"time"



	"github.com/jinzhu/gorm"
	"github.com/badoux/checkmail"


	
)



//User model for User entity
type User struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	FirstName  string   `gorm:"size:255;not null;" json:"firstname"`
	LastName  string    `gorm:"size:255;not null;" json:"lastname"`
	Email      string    `gorm:"size:100;not null;unique_index" json:"email"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}



//BeforeSave function to trim white space 
//from
func (u *User) BeforeSave() {
	//u.ID = 0
	u.FirstName = html.EscapeString(strings.TrimSpace(u.FirstName))
	u.LastName = html.EscapeString(strings.TrimSpace(u.LastName))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}


//Validate user input from mobile client
func (u *User) Validate(str string) error {
	switch strings.ToLower(str) {
	case "register":
		if u.FirstName == "" {
			return errors.New("First name is required")
		}
		if u.LastName == "" {
			return errors.New("Last name is required Password")
		}
		if u.Email == "" {
			return errors.New("Email is required")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}

		return nil
	case "login":
		if u.Email == "" {
			return errors.New("Email is required")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil

	default:
		if u.FirstName == "" {
			return errors.New("First name is required")
		}
		if u.LastName == "" {
			return errors.New("Last name is required Password")
		}
		if u.Email == "" {
			return errors.New("Email is required")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil
	}
}



//SaveUser new user 
func (u *User) SaveUser(db *gorm.DB) (*User, error) {

//if Email exists dont insert new record
//return empty user struct
	  user := &User{}
	  err := db.Debug().Table("users").Where("email = ?", u.Email).First(user).Error
	
	 if user.Email != ""{
		return &User{}, err
	 }

	db.Debug().Create(&u)

	//generate Token here
	
	return u, nil
	 
	 	 
	// return u, nil
}
