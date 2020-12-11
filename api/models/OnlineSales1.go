package models

import (
	"encoding/json"
	"time"

	"github.com/jinzhu/gorm"
)

type Onlinesales1 struct {
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

func ShowPaymentMethodVAOnlineSales(db *gorm.DB) (error, []Onlinesales1) {
	var datas []Onlinesales1
	query := `SELECT * FROM onlinesales WHERE payment_method LIKE '%Virtual Account%'`
	err := db.Raw(query).Scan(&datas).Error
	if err != nil {
		return err, nil
	}

	return nil, datas
}
