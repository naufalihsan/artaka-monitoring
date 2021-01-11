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
	ID            uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Create_dtm    time.Time `json:"create_dtm"`
	User_id       string    `json:"user_id"`
	Email         string    `json:"email"`
	Owner_name    string    `json:"owner_name"`
	Referral_code string    `json:"referral_code"`
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
func ShowSubscribers(db *gorm.DB) (error, []*MerchantsData) {
	query := `select id, create_dtm, user_id, email, owner_name, referral_code from subscribers`
	var merchant []*MerchantsData
	err := db.Raw(query).Scan(&merchant).Error
	if err != nil {
		return err, nil
	}
	return nil, merchant
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
