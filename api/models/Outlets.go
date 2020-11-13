package models

import (
	"encoding/json"
	"time"
)

type Outlets struct {
	ID                     uint32          `gorm:"primary_key;auto_increment" json:"id"`
	Create_dtm             time.Time       `json:"create_dtm"`
	User_id                string          `json:"user_id"`
	Outlet_id              string          `json:"outlet_id"`
	Nama                   string          `json:"nama"`
	Address                string          `json:"address"`
	Phone                  string          `json:"phone"`
	Business_category      string          `json:"business_category"`
	Is_active              string          `json:"is_active"`
	Accounts               json.RawMessage `json:"accounts"`
	Images                 json.RawMessage `json:"images"`
	Mini_website_url       string          `json:"mini_website_url"`
	Is_online_store_active string          `json:"is_online_store_active"`
}
