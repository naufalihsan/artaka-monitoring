package models

import (
	"errors"
	"html"
	"log"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/gunturbudikurniawan/Artaka/api/security"
	"github.com/jinzhu/gorm"
)

// Admin struct
type Admin struct {
	ID              uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Phone           string    `gorm:"size:100;" json:"phone"`
	Username        string    `gorm:"size:255;not null;unique" json:"username"`
	Create_dtm      time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"create_dtm`
	Email           string    `gorm:"size:100;not null;unique" json:"email"`
	Secret_password string    `json:"secret_password"`
	Role            string    `gorm:json:"role"`
	Referral_code   string    `gorm: json:"referral_code"`
}

func (a *Admin) BeforeSave() error {
	hashedPassword, err := security.Hash(a.Secret_password)
	if err != nil {
		return err
	}
	a.Secret_password = string(hashedPassword)
	return nil
}

func (a *Admin) Prepare() {
	a.Username = html.EscapeString(strings.TrimSpace(a.Username))
	a.Email = html.EscapeString(strings.TrimSpace(a.Email))
	a.Phone = html.EscapeString(strings.TrimSpace(a.Phone))
	a.Create_dtm = time.Now()

}

func (a *Admin) Validate(action string) map[string]string {
	var errorMessages = make(map[string]string)
	var err error

	switch strings.ToLower(action) {
	case "update":
		if a.Email == "" {
			err = errors.New("Required Email")
			errorMessages["Required_email"] = err.Error()
		}
		if a.Email != "" {
			if err = checkmail.ValidateFormat(a.Email); err != nil {
				err = errors.New("Invalid Email")
				errorMessages["Invalid_email"] = err.Error()
			}
		}

	case "login":
		if a.Secret_password == "" {
			err = errors.New("Required Password")
			errorMessages["Required_password"] = err.Error()
		}
		if a.Email == "" && a.Phone == "" {
			err = errors.New("Required Email")
			errorMessages["Required_email"] = err.Error()
		}
		if a.Email != "" {
			if err = checkmail.ValidateFormat(a.Email); err != nil {
				err = errors.New("Invalid Email")
				errorMessages["Invalid_email"] = err.Error()
			}
		}
	case "forgotpassword":
		if a.Email == "" {
			err = errors.New("Required Email")
			errorMessages["Required_email"] = err.Error()
		}
		if a.Email != "" {
			if err = checkmail.ValidateFormat(a.Email); err != nil {
				err = errors.New("Invalid Email")
				errorMessages["Invalid_email"] = err.Error()
			}
		}
	default:
		if a.Username == "" {
			err = errors.New("Required Username")
			errorMessages["Required_username"] = err.Error()
		}
		if a.Secret_password == "" {
			err = errors.New("Required Password")
			errorMessages["Required_password"] = err.Error()
		}
		if a.Secret_password != "" && len(a.Secret_password) < 6 {
			err = errors.New("Password should be atleast 6 characters")
			errorMessages["Invalid_password"] = err.Error()
		}
		if a.Email == "" {
			err = errors.New("Required Email")
			errorMessages["Required_email"] = err.Error()

		}
		if a.Email != "" {
			if err = checkmail.ValidateFormat(a.Email); err != nil {
				err = errors.New("Invalid Email")
				errorMessages["Invalid_email"] = err.Error()
			}
		}
	}
	return errorMessages
}

func (a *Admin) SaveAdmin(db *gorm.DB) (*Admin, error) {
	var err error
	err = db.Debug().Create(&a).Error
	if err != nil {
		return &Admin{}, err
	}

	return a, nil
}

func (a *Admin) UpdateAdmin(db *gorm.DB, uid uint32) (*Admin, error) {

	if a.Secret_password != "" {
		err := a.BeforeSave()
		if err != nil {
			log.Fatal(err)
		}

		db = db.Debug().Model(&Admin{}).Where("id = ?", uid).Take(&Admin{}).UpdateColumns(
			map[string]interface{}{
				"secret_password": a.Secret_password,
				"email":           a.Email,
				"create_dtm":      time.Now(),
			},
		)
	}
	db = db.Debug().Model(&Admin{}).Where("id = ?", uid).Take(&Admin{}).UpdateColumns(
		map[string]interface{}{
			"email":      a.Email,
			"create_dtm": time.Now(),
		},
	)
	if db.Error != nil {
		return &Admin{}, db.Error
	}

	err := db.Debug().Model(&Admin{}).Where("id = ?", uid).Take(&a).Error
	if err != nil {
		return &Admin{}, err
	}
	return a, nil
}
