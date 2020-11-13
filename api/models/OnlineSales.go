package models

import (
	"encoding/json"
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Onlinesales struct {
	ID               uint32          `gorm:"primary_key;auto_increment" json:"id"`
	Create_dtm       time.Time       `json:"create_dtm"`
	Sales_id         string          `json:"sales_id"`
	User_id          string          `json:"user_id"`
	Outlet_id        string          `json:"outlet_id"`
	Customer_id      string          `json:"customer_id"`
	Customer         json.RawMessage `json:"customer"`
	Products         json.RawMessage `json:"products"`
	Subtotal         int             `json:"subtotal"`
	Total_diskon     int             `json:"total_diskon"`
	Total_tax        json.RawMessage `json:"total_tax"`
	Total_bill       int             `json:"total_bill"`
	Payment_method   string          `json:"payment_method"`
	Payment_account  string          `json:"payment_account"`
	Payment_due_date string          `json:"payment_due_date"`
	Total_payment    int             `json:"total_payment"`
	Expedition       string          `json:"expedition"`
	Service          string          `json:"service"`
	Weight           int             `json:"weight"`
	Delivery_cost    int             `json:"delivery_cost"`
	Notes            string          `json:"notes"`
	Total_buy_cost   int             `json:"total_buy_cost"`
	Payment_date     string          `json:"payment_date"`
	Reward_id        string          `json:"reward_id"`
	Points_redeem    int             `json:"points_redeem"`
	Order_status     string          `json:"order_status"`
	Shipment_number  string          `json:"shipment_number"`
}

func (k *Onlinesales) Prepare() {

	k.Payment_account = html.EscapeString(strings.TrimSpace(k.Payment_account))
}

func (k *Onlinesales) Validate() map[string]string {

	var err error

	var errorMessages = make(map[string]string)

	if k.Customer_id == "" {
		err = errors.New("Required Customer")
		errorMessages["Required_Customer"] = err.Error()
	}
	if k.Outlet_id == "" {
		err = errors.New("Required Outlet")
		errorMessages["Required_Outlet"] = err.Error()
	}
	return errorMessages
}

func (k *Onlinesales) SaveOnlineSales(db *gorm.DB) (*Onlinesales, error) {
	var err error
	err = db.Debug().Model(&Onlinesales{}).Create(&k).Error
	if err != nil {
		return &Onlinesales{}, err
	}

	if k.User_id == "" {
		err = db.Debug().Model(&Onlinesales{}).Where("user_id = ?", k.User_id).Error
		if err != nil {
			return &Onlinesales{}, err
		}
	}
	return k, nil
}

func (k *Onlinesales) FindOnline(db *gorm.DB) (*Onlinesales, error) {
	var err error
	err = db.Debug().Model(Onlinesales{}).Where("created_at > ?", time.Now().Add(-24*time.Hour)).Take(&k).Error
	if err != nil {
		return &Onlinesales{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Onlinesales{}, errors.New("User Not Found")
	}
	return k, err
}
