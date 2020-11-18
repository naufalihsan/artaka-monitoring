package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Sales struct {
	ID             uint32          `gorm:"primary_key;auto_increment" json:"id"`
	CreateDTM      time.Time       `json:"create_dtm"`
	SalesID        string          `json:"sales_id"`
	UserID         string          `json:"user_id"`
	OutletID       string          `json:"outlet_id"`
	SalesType      string          `json:"sales_type"`
	CustomerID     string          `json:"customer_id"`
	Products       json.RawMessage `json:"products"`
	Subtotal       int             `json:"subtotal"`
	TotalDiskon    int             `json:"total_diskon"`
	TotalTax       json.RawMessage `json:"total_tax"`
	TotalBill      int             `json:"total_bill"`
	PaymentMethod  string          `json:"payment_method"`
	PaymentDueDate string          `json:"payment_due_date"`
	TotalPayment   int             `json:"total_payment"`
	Exchange       int             `json:"exchange"`
	Notes          string          `json:"notes"`
	TotalBuyCost   int             `json:"total_buy_cost"`
	PaymentDate    string          `json:"payment_date"`
	RewardID       string          `json:"Reward_id"`
	PointsRedeem   int             `json:"points_redeem"`
}
type Data struct {
	UserID            string
	OwnerName         string
	Email             string
	Create_dtm        string
	Toko_name_address string
	Feedback          string
}

func (w *Sales) Prepare() {
	w.SalesType = html.EscapeString(strings.TrimSpace(w.SalesType))
	w.CreateDTM = time.Now()
}

func (w *Sales) Validate() map[string]string {

	var err error

	var errorMessages = make(map[string]string)

	if w.CustomerID == "" {
		err = errors.New("Required Customer")
		errorMessages["Required_Customer"] = err.Error()
	}
	if w.OutletID == "" {
		err = errors.New("Required Outlet")
		errorMessages["Required_Outlet"] = err.Error()
	}
	return errorMessages
}

func (w *Sales) SaveSales(db *gorm.DB) (*Sales, error) {
	var err error
	err = db.Debug().Model(&Sales{}).Create(&w).Error
	if err != nil {
		return &Sales{}, err
	}
	if w.UserID == "" {
		err = db.Debug().Model(&Sales{}).Where("user_id = ?", w.UserID).Error
		if err != nil {
			return &Sales{}, err
		}
	}
	return w, nil
}

func (w *Sales) FindSales(db *gorm.DB) (*Sales, error) {
	var err error
	err = db.Debug().Model(Sales{}).Where("created_at > ?", time.Now().Add(-168*time.Hour)).Take(&w).Error
	if err != nil {
		return &Sales{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Sales{}, errors.New("User Not Found")
	}
	return w, err
}

// heroku pg:psql --app artaka < /home/guntur/dbase.sql
func Show(db *gorm.DB) (error, []Data) {
	var datas []Data

	query := `select Z.user_id, Z.owner_name, Z.email, MAX(Z.create_dtm) as create_dtm, Z.toko_name_address,Z.feedback from (
		select user_id, owner_name, email, (select create_dtm from sales where user_id = b.user_id order by id desc limit 1), (select concat(nama,'|', address) as nama from outlets where user_id = b.user_id limit 1) as toko_name_address, (select concat(content,'|', updated_at) as content from posts where content IS NOT NULL AND phone = b.user_id limit 1) as feedback from subscribers b
		UNION
		select user_id, owner_name, email, (select create_dtm from onlinesales where user_id = b.user_id order by id desc limit 1), (select concat(nama,'|', address) as nama from outlets where user_id = b.user_id limit 1) as toko_name_address, (select concat(content,'|', updated_at) as content from posts where content IS NOT NULL AND  phone = b.user_id limit 1) as feedback from subscribers b
		UNION
		select user_id, owner_name, email, (select create_dtm from saved_orders so where user_id = b.user_id order by id desc limit 1), (select concat(nama,'|', address) as nama from outlets where user_id = b.user_id limit 1) as toko_name_address, (select concat(content,'|', updated_at) as content from posts where content IS NOT NULL AND  phone = b.user_id limit 1) as feedback  from subscribers b
		) as Z GROUP BY Z.user_id, Z.owner_name, Z.email, Z.toko_name_address,Z.feedback`
	err := db.Raw(query).Scan(&datas).Error
	if err != nil {
		fmt.Println(err)
		return err, nil
	}
	var res []Data
	for i := 0; i < len(datas); i++ {
		if datas[i].Feedback == "" {
			res = append(res, datas[i])
		}
	}
	return nil, res
}
func Show1(db *gorm.DB) (error, []Data) {
	var datas []Data

	query := `select Z.user_id, Z.owner_name, Z.email, MAX(Z.create_dtm) as create_dtm, Z.toko_name_address,Z.feedback from (
		select user_id, owner_name, email, (select create_dtm from sales where user_id = b.user_id order by id desc limit 1), (select concat(nama,'|', address) as nama from outlets where user_id = b.user_id limit 1) as toko_name_address, (select concat(content,'|', updated_at) as content from posts where content IS NOT NULL AND phone = b.user_id limit 1) as feedback from subscribers b
		UNION
		select user_id, owner_name, email, (select create_dtm from onlinesales where user_id = b.user_id order by id desc limit 1), (select concat(nama,'|', address) as nama from outlets where user_id = b.user_id limit 1) as toko_name_address, (select concat(content,'|', updated_at) as content from posts where content IS NOT NULL AND  phone = b.user_id limit 1) as feedback from subscribers b
		UNION
		select user_id, owner_name, email, (select create_dtm from saved_orders so where user_id = b.user_id order by id desc limit 1), (select concat(nama,'|', address) as nama from outlets where user_id = b.user_id limit 1) as toko_name_address, (select concat(content,'|', updated_at) as content from posts where content IS NOT NULL AND  phone = b.user_id limit 1) as feedback  from subscribers b
		) as Z GROUP BY Z.user_id, Z.owner_name, Z.email, Z.toko_name_address,Z.feedback`
	err := db.Raw(query).Scan(&datas).Error
	if err != nil {
		fmt.Println(err)
		return err, nil
	}
	var res []Data
	for i := 0; i < len(datas); i++ {
		if datas[i].Feedback != "" {
			res = append(res, datas[i])
		}
	}
	return nil, res
}
func Allshow(db *gorm.DB) (error, []Data) {
	var datas []Data

	query := `select Z.user_id, Z.owner_name, Z.email, MAX(Z.create_dtm) as create_dtm, Z.toko_name_address,Z.feedback from (
		select user_id, owner_name, email, (select create_dtm from sales where user_id = b.user_id order by id desc limit 1), (select concat(nama,'|', address) as nama from outlets where user_id = b.user_id limit 1) as toko_name_address, (select concat(content,'|', updated_at) as content from posts where content IS NOT NULL AND phone = b.user_id limit 1) as feedback from subscribers b
		UNION
		select user_id, owner_name, email, (select create_dtm from onlinesales where user_id = b.user_id order by id desc limit 1), (select concat(nama,'|', address) as nama from outlets where user_id = b.user_id limit 1) as toko_name_address, (select concat(content,'|', updated_at) as content from posts where content IS NOT NULL AND  phone = b.user_id limit 1) as feedback from subscribers b
		UNION
		select user_id, owner_name, email, (select create_dtm from saved_orders so where user_id = b.user_id order by id desc limit 1), (select concat(nama,'|', address) as nama from outlets where user_id = b.user_id limit 1) as toko_name_address, (select concat(content,'|', updated_at) as content from posts where content IS NOT NULL AND  phone = b.user_id limit 1) as feedback  from subscribers b
		) as Z GROUP BY Z.user_id, Z.owner_name, Z.email, Z.toko_name_address,Z.feedback`
	err := db.Raw(query).Scan(&datas).Error
	if err != nil {
		fmt.Println(err)
		return err, nil
	}
	return nil, datas
}
