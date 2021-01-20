package models

import (
	"encoding/json"
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Saved_orders struct {
	ID              uint            `gorm:"primary_key;auto_increment" json:"id"`
	Create_dtm      time.Time       `json:"create_dtm"`
	User_id         string          `json:"user_id"`
	Outlet_id       string          `json:"outlet_id"`
	Saved_orders_id string          `json:"saved_orders_id"`
	Name            string          `json:"name"`
	Phone           string          `json:"phone"`
	Orders          json.RawMessage `json:"orders"`
	Table_id        string          `json:"table_id"`
}

func (s *Saved_orders) Prepare() {
	s.Name = html.EscapeString(strings.TrimSpace(s.Name))
}

func (s *Saved_orders) Validate() map[string]string {

	var err error

	var errorMessages = make(map[string]string)

	if s.Name == "" {
		err = errors.New("Required Name")
		errorMessages["Required_Name"] = err.Error()

	}
	if s.Outlet_id == "" {
		err = errors.New("Required Content")
		errorMessages["Required_content"] = err.Error()

	}
	return errorMessages
}

func (s *Saved_orders) SaveOrder(db *gorm.DB) (*Saved_orders, error) {
	var err error
	err = db.Debug().Model(&Saved_orders{}).Create(&s).Error
	if err != nil {
		return &Saved_orders{}, err
	}
	if s.User_id == "" {
		err = db.Debug().Model(&Subscribers{}).Where("user_id = ?", s.User_id).Error
		if err != nil {
			return &Saved_orders{}, err
		}
	}
	return s, nil
}

func (s *Saved_orders) FindSaved(db *gorm.DB) (*Saved_orders, error) {
	var err error
	err = db.Debug().Model(Saved_orders{}).Where("created_at > ?", time.Now().Add(-24*time.Hour)).Take(&s).Error
	if err != nil {
		return &Saved_orders{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Saved_orders{}, errors.New("User Not Found")
	}
	return s, err
}
