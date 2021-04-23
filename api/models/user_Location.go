package models

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

//UserLocation is the struct for
//the admin who will register user's location
//for either visitors etc to
//scan qrcode to take attendance
type UserLocation struct {
	LocationID uint32    `gorm:"primary_key;auto_increment" json:"location_id"`
	Latitude   float64   `gorm:"size:255;not null;" json:"latitude"`
	Longitude  float64   `gorm:"size:255;not null;" json:"longitude"`
	Admin      User      `json:"author"`
	AdminID    uint32    `sql:"type:int REFERENCES users(id)" json:"admin_id"`
	Title      string    `gorm:"size:255;not null;unique" json:"title"`
	Content    string    `gorm:"size:255;not null;" json:"content"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

//BeforeSave func trims inputs before saving in DB
//ul ->UserLocation
func (ul *UserLocation) BeforeSave() {
	ul.LocationID = 0
	//ul.Latitude 	= html.EscapeString(strings.TrimSpace(ul.Latitude))
	//ul.Longitude	= html.EscapeString(strings.TrimSpace(p.Content))
	ul.CreatedAt = time.Now()
	ul.UpdatedAt = time.Now()
}

//ValidateData before Admin can create 
//new checkin group
func (ul *UserLocation) ValidateData() error {

	if ul.Title == "" {
		return errors.New("Required Title")
	}
	
	if ul.AdminID < 1 {
		return errors.New("Required Admin")
	}
	return nil
}


//SaveAminSelectedLocation is a func to save the Admin's 
//selected location where qrcode will
//scanned for attendance
func (ul *UserLocation) SaveAminSelectedLocation(db *gorm.DB) (*UserLocation, error) {
	var err error
	err = db.Debug().Model(&UserLocation{}).Create(&ul).Error
	if err != nil {
		return &UserLocation{}, err
	}
	if ul.LocationID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", ul.AdminID).Take(&ul.Admin).Error
		if err != nil {
			return &UserLocation{}, err
		}
	}
	return ul, nil
}