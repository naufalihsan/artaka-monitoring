package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Post struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Phone     string    `gorm:"size:255;not null;" json:"phone"`
	Content   string    `gorm:"text;not null;" json:"content"`
	Author    Admin     `json:"author"`
	Boolean   bool      `json:"boolean" gorm:"type:boolean; default:'0'"`
	AuthorID  uint32    `gorm:"not null" json:"author_id"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (p *Post) Prepare() {
	p.Phone = html.EscapeString(strings.TrimSpace(p.Phone))
	p.Content = html.EscapeString(strings.TrimSpace(p.Content))
	p.Author = Admin{}
	p.UpdatedAt = time.Now()
}

func (p *Post) Validate() map[string]string {

	var err error

	var errorMessages = make(map[string]string)

	if p.Phone == "" {
		err = errors.New("Required Phone")
		errorMessages["Required_Phone"] = err.Error()

	}

	if p.AuthorID < 1 {
		err = errors.New("Required Author")
		errorMessages["Required_author"] = err.Error()
	}
	return errorMessages
}

func (p *Post) SavePost(db *gorm.DB) (*Post, error) {
	var err error
	err = db.Debug().Model(&Post{}).Create(&p).Error
	if err != nil {
		return &Post{}, err
	}
	if p.ID != 0 {
		err = db.Debug().Model(&Admin{}).Where("id = ?", p.AuthorID).Take(&p.Author).Error
		if err != nil {
			return &Post{}, err
		}
	}
	return p, nil
}

func (p *Post) FindPostByID(db *gorm.DB, pid uint64) (*Post, error) {
	var err error
	err = db.Debug().Model(&Post{}).Where("id = ?", pid).Take(&p).Error
	if err != nil {
		return &Post{}, err
	}
	if p.ID != 0 {
		err = db.Debug().Model(&Admin{}).Where("id = ?", p.AuthorID).Take(&p.Author).Error
		if err != nil {
			return &Post{}, err
		}
	}
	return p, nil
}

func (p *Post) DeleteAPost(db *gorm.DB) (int64, error) {

	db = db.Debug().Model(&Post{}).Where("id = ?", p.ID).Take(&Post{}).Delete(&Post{})
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
func (p *Post) UpdateAPost(db *gorm.DB) (*Post, error) {

	var err error

	if p.Boolean {
		db.Debug().Model(&Post{}).Where("id = ?",
			p.ID).Updates(Post{Phone: p.Phone, Content: p.Content, Boolean: p.Boolean, UpdatedAt: time.Now()})
	} else {
		db.Debug().Model(&Post{}).Where("id = ?",
			p.ID).Updates(Post{Phone: p.Phone, Content: p.Content, Boolean: p.Boolean, UpdatedAt: time.Now()})
	}
	if err != nil {
		return &Post{}, err
	}
	if p.ID != 0 {
		err = db.Debug().Model(&Admin{}).Where("id = ?", p.AuthorID).Take(&p.Author).Error
		if err != nil {
			return &Post{}, err
		}
	}
	// p := &Post{}
	// m := structs.Map(p)

	return p, nil
}

func (p *Post) FindUserPosts(db *gorm.DB, uid uint32) (*[]Post, error) {

	var err error
	posts := []Post{}
	err = db.Debug().Model(&Post{}).Where("author_id = ?", uid).Limit(100).Order("created_at desc").Find(&posts).Error
	if err != nil {
		return &[]Post{}, err
	}
	if len(posts) > 0 {
		for i, _ := range posts {
			err := db.Debug().Model(&Admin{}).Where("id = ?", posts[i].AuthorID).Take(&posts[i].Author).Error
			if err != nil {
				return &[]Post{}, err
			}
		}
	}
	return &posts, nil
}

//When a user is deleted, we also delete the post that the user had
func (c *Post) DeleteUserPosts(db *gorm.DB, uid uint32) (int64, error) {
	posts := []Post{}
	db = db.Debug().Model(&Post{}).Where("author_id = ?", uid).Find(&posts).Delete(&posts)
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
