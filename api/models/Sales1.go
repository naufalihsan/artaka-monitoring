package models

import (
	"encoding/json"
	"time"

	"github.com/jinzhu/gorm"
)

type Sales1 struct {
	ID               uint32          `gorm:"primary_key;auto_increment" json:"id"`
	Create_dtm       time.Time       `json:"create_dtm"`
	Sales_id         string          `json:"sales_id"`
	User_id          string          `json:"user_id"`
	Outlet_id        string          `json:"outlet_id"`
	Sales_type       string          `json:"sales_type"`
	Customer_id      string          `json:"customer_id"`
	Products         json.RawMessage `json:"products"`
	Subtotal         int             `json:"subtotal"`
	Total_diskon     int             `json:"total_diskon"`
	Total_tax        json.RawMessage `json:"total_tax"`
	Total_bill       int             `json:"total_bill"`
	Payment_method   string          `json:"payment_method"`
	Payment_due_date string          `json:"payment_due_date"`
	Total_payment    int             `json:"total_payment"`
	Exchange         int             `json:"exchange"`
	Notes            string          `json:"notes"`
	Total_buy_cost   int             `json:"total_buy_cost"`
	Payment_date     string          `json:"payment_date"`
	Reward_id        string          `json:"Reward_id"`
	Points_redeem    int             `json:"points_redeem"`
}

func ShowPaymentMethodSales(db *gorm.DB) (error, []Sales1) {
	var datas []Sales1
	query := `SELECT * FROM sales WHERE payment_method='Tunai' OR payment_method='Uang Pas'`
	err := db.Raw(query).Scan(&datas).Error
	if err != nil {
		return err, nil
	}

	return nil, datas
}
