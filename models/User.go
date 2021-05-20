package models

import (
	"errors"
	"go-blog-api/security"
	"html"

	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type User struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Nickname  string    `gorm:"size:20;not null;unique" json:"nickname"`
	Email     string    `gorm:"size:50;not null;unique" json:"email"`
	Password  string    `gorm:"size:60;not null" json:"password"`
	CreatedAt time.Time `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:current_timestamp()" json:"updated_at"`
	Posts     []Post    `gorm:"foreignKey:AuthorID; " json:"posts"`
}

func (u *User) BeforeSave() error {
	hashedPassword, err := security.Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) Prepare() {
	u.ID = 0
	u.Nickname = html.EscapeString(strings.TrimSpace(u.Nickname))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *User) Validate(action string) error {
	switch action {
	case "update":
		if u.Nickname == "" {
			return errors.New("required nickname")
		}
		if u.Password == "" {
			return errors.New("required password")
		}
		if u.Email == "" {
			return errors.New("require email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("invalid email")
		}
		return nil
	case "login":
		if u.Email == "" {
			return errors.New("require email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("invalid email")
		}
		if u.Password == "" {
			return errors.New("required password")
		}
		return nil
	default:
		if u.Nickname == "" {
			return errors.New("required nickname")
		}
		if u.Password == "" {
			return errors.New("required password")
		}
		if u.Email == "" {
			return errors.New("require email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("invalid email")
		}
		return nil
	}
}
