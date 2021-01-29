package models

import (
	"encoding/json"
	"errors"
	"html"
	"log"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/gunturbudikurniawan/Artaka/api/security"
	"github.com/jinzhu/gorm"
)

type Subscribers struct {
	ID               uint32          `gorm:"primary_key;auto_increment" json:"id"`
	Create_dtm       time.Time       `json:"create_dtm"`
	User_id          string          `json:"user_id"`
	Email            string          `json:"email"`
	Owner_name       string          `json:"owner_name"`
	Secret_password  string          `json:"secret_password"`
	Fcm_token        string          `json:"fcm_token"`
	Idcard_name      string          `json:"idcard_name"`
	Idcard_number    string          `json:"idcard_number"`
	Bank_holder_name string          `json:"bank_holder_name"`
	Bank_name        string          `json:"bank_name"`
	Bank_account     string          `json:"bank_account"`
	Idcard_image     json.RawMessage `json:"idcard_image"`
	Referral_code    string          `json:"referral_code"`
}
type MerchantsData struct {
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
	Business_category string          `json:"business_category"`
}

func (m *Subscribers) BeforeSave() error {
	hashedPassword, err := security.Hash(m.Secret_password)
	if err != nil {
		return err
	}
	m.Secret_password = string(hashedPassword)
	return nil
}

func (m *Subscribers) Prepare() {

	m.Email = html.EscapeString(strings.TrimSpace(m.Email))
	m.Fcm_token = html.EscapeString(strings.TrimSpace(m.Fcm_token))
	m.Idcard_name = html.EscapeString(strings.TrimSpace(m.Idcard_name))
	m.Idcard_number = html.EscapeString(strings.TrimSpace(m.Idcard_number))
	m.Owner_name = html.EscapeString(strings.TrimSpace(m.Owner_name))
	m.Bank_account = html.EscapeString(strings.TrimSpace(m.Bank_account))
	m.Create_dtm = time.Now()
}
func (m *Subscribers) Validate(action string) map[string]string {
	var errorMessages = make(map[string]string)
	var err error

	switch strings.ToLower(action) {
	case "update":
		if m.Email == "" {
			err = errors.New("Required Email")
			errorMessages["Required_email"] = err.Error()
			return errorMessages
		}
		if m.Email != "" {
			if err = checkmail.ValidateFormat(m.Email); err != nil {
				err = errors.New("Invalid Email")
				errorMessages["Invalid_email"] = err.Error()
				return errorMessages

			}
		}

	case "login":
		if m.Secret_password == "" {
			err = errors.New("Required Password")
			errorMessages["Required_password"] = err.Error()
			return errorMessages

		}
		if m.Email == "" {
			err = errors.New("Required Email")
			errorMessages["Required_email"] = err.Error()
			return errorMessages

		}
		if m.Email != "" {
			if err = checkmail.ValidateFormat(m.Email); err != nil {
				err = errors.New("Invalid Email")
				errorMessages["Invalid_email"] = err.Error()
				return errorMessages

			}
		}
	case "forgotpassword":
		if m.Email == "" {
			err = errors.New("Required Email")
			errorMessages["Required_email"] = err.Error()
			return errorMessages

		}
		if m.Email != "" {
			if err = checkmail.ValidateFormat(m.Email); err != nil {
				err = errors.New("Invalid Email")
				errorMessages["Invalid_email"] = err.Error()
				return errorMessages

			}
		}
	default:
		if m.Owner_name == "" {
			err = errors.New("Required Owner Name")
			errorMessages["Required Owner Name"] = err.Error()
			return errorMessages

		}
		if m.Secret_password == "" {
			err = errors.New("Required Password")
			errorMessages["Required_password"] = err.Error()
			return errorMessages

		}
		if m.Secret_password != "" && len(m.Secret_password) < 6 {
			err = errors.New("Password should be atleast 6 characters")
			errorMessages["Invalid_password"] = err.Error()
			return errorMessages

		}
		if m.Email == "" {
			err = errors.New("Required Email")
			errorMessages["Required_email"] = err.Error()
			return errorMessages

		}
		if m.Email != "" {
			if err = checkmail.ValidateFormat(m.Email); err != nil {
				err = errors.New("Invalid Email")
				errorMessages["Invalid_email"] = err.Error()
				return errorMessages

			}
		}
	}
	return errorMessages
}
func (m *Subscribers) SaveUser(db *gorm.DB) (*Subscribers, error) {

	var err error
	err = db.Debug().Create(&m).Error
	if err != nil {
		return &Subscribers{}, err
	}
	return m, nil
}

func (m *Subscribers) UpdateMerchant(db *gorm.DB, uid uint32) (*Subscribers, error) {

	if m.Secret_password != "" {
		err := m.BeforeSave()
		if err != nil {
			log.Fatal(err)
		}

		db = db.Debug().Model(&Subscribers{}).Where("id = ?", uid).Take(&Subscribers{}).UpdateColumns(
			map[string]interface{}{
				"password": m.Secret_password,
				"email":    m.Email,
			},
		)
	}
	db = db.Debug().Model(&Subscribers{}).Where("id = ?", uid).Take(&Subscribers{}).UpdateColumns(
		map[string]interface{}{
			"email": m.Email,
		},
	)
	if db.Error != nil {
		return &Subscribers{}, db.Error
	}

	err := db.Debug().Model(&Subscribers{}).Where("id = ?", uid).Take(&m).Error
	if err != nil {
		return &Subscribers{}, err
	}
	return m, nil
}

func (m *Subscribers) FindAllMerchants(db *gorm.DB) (*[]Subscribers, error) {
	var err error
	merchants := []Subscribers{}
	err = db.Debug().Model(&Subscribers{}).Limit(100).Find(&merchants).Error
	if err != nil {
		return &[]Subscribers{}, err
	}
	return &merchants, err
}
func ShowSubscribers(db *gorm.DB, referral_code string, role string) (error, []MerchantsData) {
	var datas []MerchantsData
	query := `select xx.user_id,(select owner_name from subscribers where user_id = xx.user_id limit 1) owner_name, 
	(select email from subscribers where user_id = xx.user_id limit 1) email, 
	(select create_dtm from subscribers where user_id = xx.user_id limit 1) register, max(xx.create_dtm) as create_dtm,
	(select concat(nama,'|', address) as Toko_name_address from outlets where user_id = xx.user_id limit 1) as Toko_name_address, 
	(select content as content from posts where content IS NOT NULL AND phone = xx.user_id limit 1) as feedback, 
	(select updated_at from posts where content IS NOT NULL AND phone = xx.user_id limit 1) as tanggal,
	(select boolean as boolean from posts where phone = xx.user_id limit 1) as boolean,
	(select id from posts where  phone = xx.user_id limit 1) as idpost, 
	(select to_jsonb(images) from outlets where user_id = xx.user_id limit 1) as images, 
	(select referral_code from subscribers where user_id = xx.user_id limit 1) as referral_code,
	(select business_category from outlets where user_id = xx.user_id limit 1) as business_category
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
	var res []MerchantsData
	for i := 0; i < len(datas); i++ {
		if role == "ADMIN" {
			if datas[i].Referral_code != "" || datas[i].Referral_code == "" {
				res = append(res, datas[i])
			}
		} else if strings.Contains(strings.ToUpper(datas[i].Referral_code), strings.ToUpper(referral_code)) {
			res = append(res, datas[i])
		}
	}
	return nil, res
}

func (m *Subscribers) FindMerchantByID(db *gorm.DB, uid uint32) (*Subscribers, error) {
	var err error
	err = db.Debug().Model(Subscribers{}).Where("id = ?", uid).Take(&m).Error
	if err != nil {
		return &Subscribers{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Subscribers{}, errors.New("User Not Found")
	}
	return m, err
}
