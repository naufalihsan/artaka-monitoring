package models

import (
	"encoding/json"
	"errors"
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
	UserID            string          `json:"user_id"`
	OwnerName         string          `json:"owner_name"`
	Email             string          `json:"email"`
	Register          string          `json:"register"`
	Create_dtm        string          `json:"create_dtm"`
	Toko_name_address string          `json:"toko_name_address"`
	Feedback          string          `json:"feedback"`
	Tanggal           string          `json:"tanggal"`
	Boolean           string          `json:"boolean" `
	Idpost            uint64          `json:"idpost"`
	Images            json.RawMessage `json:"images"`
	Referral_code     string          `json:"referral_code"`
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

func Show(db *gorm.DB) (error, []Data) {
	var datas []Data

	query := `select xx.user_id,(select owner_name from subscribers where user_id = xx.user_id limit 1) owner_name, 
	(select email from subscribers where user_id = xx.user_id limit 1) email, 
	(select create_dtm from subscribers where user_id = xx.user_id limit 1) register, max(xx.create_dtm) as create_dtm,
	(select concat(nama,'|', address) as Toko_name_address from outlets where user_id = xx.user_id limit 1) as Toko_name_address, 
	(select content as content from posts where content IS NOT NULL AND phone = xx.user_id limit 1) as feedback, 
	(select updated_at from posts where content IS NOT NULL AND phone = xx.user_id limit 1) as tanggal,
	(select boolean as boolean from posts where phone = xx.user_id limit 1) as boolean,
	(select id from posts where  phone = xx.user_id limit 1) as idpost, 
	(select to_jsonb(images) from outlets where user_id = xx.user_id limit 1) as images, 
	(select referral_code from subscribers where user_id = xx.user_id limit 1) as referral_code
	from(select y.user_id, max(y.create_dtm) as create_dtm from(select a.user_id, 
	(select s.create_dtm from sales s where user_id = a.user_id and create_dtm < (current_date -7) order by create_dtm desc limit 1)
	from subscribers a UNION select b.user_id, 
	(select create_dtm from onlinesales where user_id = b.user_id and create_dtm < (current_date -7) order by create_dtm desc limit 1)
	from subscribers b UNION select c.user_id, 
	(select create_dtm from saved_orders where user_id = c.user_id and create_dtm < (current_date -7) order by create_dtm desc limit 1) 
	from subscribers c) y group by y.user_id) xx where xx.user_id not in
	(select yy.user_id from
	(select y.user_id, max(y.create_dtm)from(select a.user_id, 
	(select s.create_dtm from sales s where user_id = a.user_id and create_dtm > (current_date -7) order by create_dtm desc limit 1)
	from subscribers a UNION select b.user_id, 
	(select create_dtm from onlinesales where user_id = b.user_id and create_dtm > (current_date -7) order by create_dtm desc limit 1)
	from subscribers b
	UNION
	select c.user_id, 
	(select create_dtm from saved_orders where user_id = c.user_id and create_dtm > (current_date -7) order by create_dtm desc limit 1) 
	from subscribers c) y where y.create_dtm is not null group by y.user_id) yy)  
	GROUP BY xx.user_id, (select owner_name from subscribers where user_id = xx.user_id limit 1), (select email from subscribers where user_id = xx.user_id limit 1), (select create_dtm from subscribers where user_id = xx.user_id limit 1), 
	(select concat(nama,'|', address) as nama from outlets where user_id = xx.user_id limit 1),(select content as content from posts where content IS NOT NULL AND phone = xx.user_id limit 1),
	(select updated_at from posts where content IS NOT NULL AND phone = xx.user_id limit 1),(select boolean as boolean from posts where phone = xx.user_id limit 1),
	(select id from posts where  phone = xx.user_id limit 1),(select to_jsonb(images) from outlets where user_id = xx.user_id limit 1)
	ORDER BY xx.user_id, (select owner_name from subscribers where user_id = xx.user_id limit 1), (select email from subscribers where user_id = xx.user_id limit 1),  (select create_dtm from subscribers where user_id = xx.user_id limit 1),
	(select concat(nama,'|', address) as nama from outlets where user_id = xx.user_id limit 1),(select content as content from posts where content IS NOT NULL AND phone = xx.user_id limit 1),
	(select updated_at from posts where content IS NOT NULL AND phone = xx.user_id limit 1),(select boolean as boolean from posts where phone = xx.user_id limit 1),
	(select id from posts where  phone = xx.user_id limit 1),(select to_jsonb(images) from outlets where user_id = xx.user_id limit 1)`
	err := db.Raw(query).Scan(&datas).Error
	if err != nil {
		return err, nil
	}
	var res []Data
	for i := 0; i < len(datas); i++ {
		if datas[i].Boolean == "" && datas[i].Feedback == "" {
			res = append(res, datas[i])
		}

	}
	return nil, res
}
func Show1(db *gorm.DB) (error, []Data) {
	var datas []Data
	query := `select xx.user_id,(select owner_name from subscribers where user_id = xx.user_id limit 1) owner_name, 
	(select email from subscribers where user_id = xx.user_id limit 1) email, 
	(select create_dtm from subscribers where user_id = xx.user_id limit 1) register, max(xx.create_dtm) as create_dtm,
	(select concat(nama,'|', address) as Toko_name_address from outlets where user_id = xx.user_id limit 1) as Toko_name_address, 
	(select content as content from posts where content IS NOT NULL AND phone = xx.user_id limit 1) as feedback, 
	(select updated_at from posts where content IS NOT NULL AND phone = xx.user_id limit 1) as tanggal,
	(select boolean as boolean from posts where phone = xx.user_id limit 1) as boolean,
	(select id from posts where  phone = xx.user_id limit 1) as idpost, 
	(select to_jsonb(images) from outlets where user_id = xx.user_id limit 1) as images, 
	(select referral_code from subscribers where user_id = xx.user_id limit 1) as referral_code
	from(select y.user_id, max(y.create_dtm) as create_dtm from(select a.user_id, 
	(select s.create_dtm from sales s where user_id = a.user_id and create_dtm < (current_date -7) order by create_dtm desc limit 1)
	from subscribers a UNION select b.user_id, 
	(select create_dtm from onlinesales where user_id = b.user_id and create_dtm < (current_date -7) order by create_dtm desc limit 1)
	from subscribers b UNION select c.user_id, 
	(select create_dtm from saved_orders where user_id = c.user_id and create_dtm < (current_date -7) order by create_dtm desc limit 1) 
	from subscribers c) y group by y.user_id) xx where xx.user_id not in
	(select yy.user_id from
	(select y.user_id, max(y.create_dtm)from(select a.user_id, 
	(select s.create_dtm from sales s where user_id = a.user_id and create_dtm > (current_date -7) order by create_dtm desc limit 1)
	from subscribers a UNION select b.user_id, 
	(select create_dtm from onlinesales where user_id = b.user_id and create_dtm > (current_date -7) order by create_dtm desc limit 1)
	from subscribers b
	UNION
	select c.user_id, 
	(select create_dtm from saved_orders where user_id = c.user_id and create_dtm > (current_date -7) order by create_dtm desc limit 1) 
	from subscribers c) y where y.create_dtm is not null group by y.user_id) yy)  
	GROUP BY xx.user_id, (select owner_name from subscribers where user_id = xx.user_id limit 1), (select email from subscribers where user_id = xx.user_id limit 1), (select create_dtm from subscribers where user_id = xx.user_id limit 1), 
	(select concat(nama,'|', address) as nama from outlets where user_id = xx.user_id limit 1),(select content as content from posts where content IS NOT NULL AND phone = xx.user_id limit 1),
	(select updated_at from posts where content IS NOT NULL AND phone = xx.user_id limit 1),(select boolean as boolean from posts where phone = xx.user_id limit 1),
	(select id from posts where  phone = xx.user_id limit 1),(select to_jsonb(images) from outlets where user_id = xx.user_id limit 1)
	ORDER BY xx.user_id, (select owner_name from subscribers where user_id = xx.user_id limit 1), (select email from subscribers where user_id = xx.user_id limit 1),  (select create_dtm from subscribers where user_id = xx.user_id limit 1),
	(select concat(nama,'|', address) as nama from outlets where user_id = xx.user_id limit 1),(select content as content from posts where content IS NOT NULL AND phone = xx.user_id limit 1),
	(select updated_at from posts where content IS NOT NULL AND phone = xx.user_id limit 1),(select boolean as boolean from posts where phone = xx.user_id limit 1),
	(select id from posts where  phone = xx.user_id limit 1),(select to_jsonb(images) from outlets where user_id = xx.user_id limit 1)`
	err := db.Raw(query).Scan(&datas).Error
	if err != nil {
		return err, nil
	}
	var res []Data
	for i := 0; i < len(datas); i++ {
		if datas[i].Feedback != "" || datas[i].Boolean != "" {
			res = append(res, datas[i])
		}
	}
	return nil, res
}
func Allshow(db *gorm.DB) (error, []Data) {
	var datas []Data
	query := `select xx.user_id,(select owner_name from subscribers where user_id = xx.user_id limit 1) owner_name, 
	(select email from subscribers where user_id = xx.user_id limit 1) email, 
	(select create_dtm from subscribers where user_id = xx.user_id limit 1) register, max(xx.create_dtm) as create_dtm,
	(select concat(nama,'|', address) as Toko_name_address from outlets where user_id = xx.user_id limit 1) as Toko_name_address, 
	(select content as content from posts where content IS NOT NULL AND phone = xx.user_id limit 1) as feedback, 
	(select updated_at from posts where content IS NOT NULL AND phone = xx.user_id limit 1) as tanggal,
	(select boolean as boolean from posts where phone = xx.user_id limit 1) as boolean,
	(select id from posts where  phone = xx.user_id limit 1) as idpost, 
	(select to_jsonb(images) from outlets where user_id = xx.user_id limit 1) as images, 
	(select referral_code from subscribers where user_id = xx.user_id limit 1) as referral_code
	from(select y.user_id, max(y.create_dtm) as create_dtm from(select a.user_id, 
	(select s.create_dtm from sales s where user_id = a.user_id and create_dtm < (current_date -7) order by create_dtm desc limit 1)
	from subscribers a UNION select b.user_id, 
	(select create_dtm from onlinesales where user_id = b.user_id and create_dtm < (current_date -7) order by create_dtm desc limit 1)
	from subscribers b UNION select c.user_id, 
	(select create_dtm from saved_orders where user_id = c.user_id and create_dtm < (current_date -7) order by create_dtm desc limit 1) 
	from subscribers c) y group by y.user_id) xx where xx.user_id not in
	(select yy.user_id from
	(select y.user_id, max(y.create_dtm)from(select a.user_id, 
	(select s.create_dtm from sales s where user_id = a.user_id and create_dtm > (current_date -7) order by create_dtm desc limit 1)
	from subscribers a UNION select b.user_id, 
	(select create_dtm from onlinesales where user_id = b.user_id and create_dtm > (current_date -7) order by create_dtm desc limit 1)
	from subscribers b
	UNION
	select c.user_id, 
	(select create_dtm from saved_orders where user_id = c.user_id and create_dtm > (current_date -7) order by create_dtm desc limit 1) 
	from subscribers c) y where y.create_dtm is not null group by y.user_id) yy)  
	GROUP BY xx.user_id, (select owner_name from subscribers where user_id = xx.user_id limit 1), (select email from subscribers where user_id = xx.user_id limit 1), (select create_dtm from subscribers where user_id = xx.user_id limit 1), 
	(select concat(nama,'|', address) as nama from outlets where user_id = xx.user_id limit 1),(select content as content from posts where content IS NOT NULL AND phone = xx.user_id limit 1),
	(select updated_at from posts where content IS NOT NULL AND phone = xx.user_id limit 1),(select boolean as boolean from posts where phone = xx.user_id limit 1),
	(select id from posts where  phone = xx.user_id limit 1),(select to_jsonb(images) from outlets where user_id = xx.user_id limit 1)
	ORDER BY xx.user_id, (select owner_name from subscribers where user_id = xx.user_id limit 1), (select email from subscribers where user_id = xx.user_id limit 1),  (select create_dtm from subscribers where user_id = xx.user_id limit 1),
	(select concat(nama,'|', address) as nama from outlets where user_id = xx.user_id limit 1),(select content as content from posts where content IS NOT NULL AND phone = xx.user_id limit 1),
	(select updated_at from posts where content IS NOT NULL AND phone = xx.user_id limit 1),(select boolean as boolean from posts where phone = xx.user_id limit 1),
	(select id from posts where  phone = xx.user_id limit 1),(select to_jsonb(images) from outlets where user_id = xx.user_id limit 1)`
	err := db.Raw(query).Scan(&datas).Error
	if err != nil {
		return err, nil
	}
	var res []Data
	for i := 0; i < len(datas); i++ {
		res = append(res, datas[i])
	}
	return nil, res
}

func NotRespon(db *gorm.DB) (error, []Data) {
	var datas []Data
	query := `select xx.user_id,(select owner_name from subscribers where user_id = xx.user_id limit 1) owner_name, 
	(select email from subscribers where user_id = xx.user_id limit 1) email, 
	(select create_dtm from subscribers where user_id = xx.user_id limit 1) register, max(xx.create_dtm) as create_dtm,
	(select concat(nama,'|', address) as Toko_name_address from outlets where user_id = xx.user_id limit 1) as Toko_name_address, 
	(select content as content from posts where content IS NOT NULL AND phone = xx.user_id limit 1) as feedback, 
	(select updated_at from posts where content IS NOT NULL AND phone = xx.user_id limit 1) as tanggal,
	(select boolean as boolean from posts where phone = xx.user_id limit 1) as boolean,
	(select id from posts where  phone = xx.user_id limit 1) as idpost, 
	(select to_jsonb(images) from outlets where user_id = xx.user_id limit 1) as images, 
	(select referral_code from subscribers where user_id = xx.user_id limit 1) as referral_code
	from(select y.user_id, max(y.create_dtm) as create_dtm from(select a.user_id, 
	(select s.create_dtm from sales s where user_id = a.user_id and create_dtm < (current_date -7) order by create_dtm desc limit 1)
	from subscribers a UNION select b.user_id, 
	(select create_dtm from onlinesales where user_id = b.user_id and create_dtm < (current_date -7) order by create_dtm desc limit 1)
	from subscribers b UNION select c.user_id, 
	(select create_dtm from saved_orders where user_id = c.user_id and create_dtm < (current_date -7) order by create_dtm desc limit 1) 
	from subscribers c) y group by y.user_id) xx where xx.user_id not in
	(select yy.user_id from
	(select y.user_id, max(y.create_dtm)from(select a.user_id, 
	(select s.create_dtm from sales s where user_id = a.user_id and create_dtm > (current_date -7) order by create_dtm desc limit 1)
	from subscribers a UNION select b.user_id, 
	(select create_dtm from onlinesales where user_id = b.user_id and create_dtm > (current_date -7) order by create_dtm desc limit 1)
	from subscribers b
	UNION
	select c.user_id, 
	(select create_dtm from saved_orders where user_id = c.user_id and create_dtm > (current_date -7) order by create_dtm desc limit 1) 
	from subscribers c) y where y.create_dtm is not null group by y.user_id) yy)  
	GROUP BY xx.user_id, (select owner_name from subscribers where user_id = xx.user_id limit 1), (select email from subscribers where user_id = xx.user_id limit 1), (select create_dtm from subscribers where user_id = xx.user_id limit 1), 
	(select concat(nama,'|', address) as nama from outlets where user_id = xx.user_id limit 1),(select content as content from posts where content IS NOT NULL AND phone = xx.user_id limit 1),
	(select updated_at from posts where content IS NOT NULL AND phone = xx.user_id limit 1),(select boolean as boolean from posts where phone = xx.user_id limit 1),
	(select id from posts where  phone = xx.user_id limit 1),(select to_jsonb(images) from outlets where user_id = xx.user_id limit 1)
	ORDER BY xx.user_id, (select owner_name from subscribers where user_id = xx.user_id limit 1), (select email from subscribers where user_id = xx.user_id limit 1),  (select create_dtm from subscribers where user_id = xx.user_id limit 1),
	(select concat(nama,'|', address) as nama from outlets where user_id = xx.user_id limit 1),(select content as content from posts where content IS NOT NULL AND phone = xx.user_id limit 1),
	(select updated_at from posts where content IS NOT NULL AND phone = xx.user_id limit 1),(select boolean as boolean from posts where phone = xx.user_id limit 1),
	(select id from posts where  phone = xx.user_id limit 1),(select to_jsonb(images) from outlets where user_id = xx.user_id limit 1)`
	err := db.Raw(query).Scan(&datas).Error
	if err != nil {
		return err, nil
	}
	var res []Data
	for i := 0; i < len(datas); i++ {
		if datas[i].Boolean != "1" {
			res = append(res, datas[i])
		}
	}
	return nil, res
}

func ResponForWiranesia(db *gorm.DB) (error, []Data) {
	var datas []Data
	query := `select xx.user_id,(select owner_name from subscribers where user_id = xx.user_id limit 1) owner_name, 
	(select email from subscribers where user_id = xx.user_id limit 1) email, 
	(select create_dtm from subscribers where user_id = xx.user_id limit 1) register, max(xx.create_dtm) as create_dtm,
	(select concat(nama,'|', address) as Toko_name_address from outlets where user_id = xx.user_id limit 1) as Toko_name_address, 
	(select content as content from posts where content IS NOT NULL AND phone = xx.user_id limit 1) as feedback, 
	(select updated_at from posts where content IS NOT NULL AND phone = xx.user_id limit 1) as tanggal,
	(select boolean as boolean from posts where phone = xx.user_id limit 1) as boolean,
	(select id from posts where  phone = xx.user_id limit 1) as idpost, 
	(select to_jsonb(images) from outlets where user_id = xx.user_id limit 1) as images, 
	(select referral_code from subscribers where user_id = xx.user_id limit 1) as referral_code
	from(select y.user_id, max(y.create_dtm) as create_dtm from(select a.user_id, 
	(select s.create_dtm from sales s where user_id = a.user_id and create_dtm < (current_date -7) order by create_dtm desc limit 1)
	from subscribers a UNION select b.user_id, 
	(select create_dtm from onlinesales where user_id = b.user_id and create_dtm < (current_date -7) order by create_dtm desc limit 1)
	from subscribers b UNION select c.user_id, 
	(select create_dtm from saved_orders where user_id = c.user_id and create_dtm < (current_date -7) order by create_dtm desc limit 1) 
	from subscribers c) y group by y.user_id) xx where xx.user_id not in
	(select yy.user_id from
	(select y.user_id, max(y.create_dtm)from(select a.user_id, 
	(select s.create_dtm from sales s where user_id = a.user_id and create_dtm > (current_date -7) order by create_dtm desc limit 1)
	from subscribers a UNION select b.user_id, 
	(select create_dtm from onlinesales where user_id = b.user_id and create_dtm > (current_date -7) order by create_dtm desc limit 1)
	from subscribers b
	UNION
	select c.user_id, 
	(select create_dtm from saved_orders where user_id = c.user_id and create_dtm > (current_date -7) order by create_dtm desc limit 1) 
	from subscribers c) y where y.create_dtm is not null group by y.user_id) yy)  
	GROUP BY xx.user_id, (select owner_name from subscribers where user_id = xx.user_id limit 1), (select email from subscribers where user_id = xx.user_id limit 1), (select create_dtm from subscribers where user_id = xx.user_id limit 1), 
	(select concat(nama,'|', address) as nama from outlets where user_id = xx.user_id limit 1),(select content as content from posts where content IS NOT NULL AND phone = xx.user_id limit 1),
	(select updated_at from posts where content IS NOT NULL AND phone = xx.user_id limit 1),(select boolean as boolean from posts where phone = xx.user_id limit 1),
	(select id from posts where  phone = xx.user_id limit 1),(select to_jsonb(images) from outlets where user_id = xx.user_id limit 1)
	ORDER BY xx.user_id, (select owner_name from subscribers where user_id = xx.user_id limit 1), (select email from subscribers where user_id = xx.user_id limit 1),  (select create_dtm from subscribers where user_id = xx.user_id limit 1),
	(select concat(nama,'|', address) as nama from outlets where user_id = xx.user_id limit 1),(select content as content from posts where content IS NOT NULL AND phone = xx.user_id limit 1),
	(select updated_at from posts where content IS NOT NULL AND phone = xx.user_id limit 1),(select boolean as boolean from posts where phone = xx.user_id limit 1),
	(select id from posts where  phone = xx.user_id limit 1),(select to_jsonb(images) from outlets where user_id = xx.user_id limit 1)`
	err := db.Raw(query).Scan(&datas).Error
	if err != nil {
		return err, nil
	}
	var res []Data
	for i := 0; i < len(datas); i++ {
		if datas[i].Referral_code == "Wiranesia" {
			res = append(res, datas[i])
		}
	}
	return nil, res
}
