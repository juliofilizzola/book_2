package models

import (
	"errors"
	"github.com/badoux/checkmail"
	"github.com/juliofilizzola/book_2/auth"
	"gorm.io/gorm"
	"strings"
)

type User struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Email       string `gorm:"unique;not null"`
	Nick        string `gorm:"unique;not null"`
	Password    string
	Followers   []User        `gorm:"many2many:follower_id" json:"followers"`
	Publication []Publication `gorm:"foreignKey:AuthId"`
}

func (u User) PrepareData(edit bool) error {
	if err := u.validationData(edit); err != nil {
		return err
	}

	if err := u.formatData(edit); err != nil {
		return err
	}

	return nil
}

func (u User) validationData(edit bool) error {
	if u.Name == "" {
		return errors.New("name has required")
	}

	if u.Nick == "" {
		return errors.New("nick has required")
	}

	if u.Email == "" {
		return errors.New("email has required")
	}

	if err := checkmail.ValidateFormat(u.Email); err != nil {
		return errors.New("email invalid")
	}

	if u.Password == "" && !edit {
		return errors.New("password has required")
	}

	return nil
}

func (u User) formatData(edit bool) error {
	u.Name = strings.TrimSpace(u.Name)
	u.Email = strings.TrimSpace(u.Email)
	u.Nick = strings.TrimSpace(u.Nick)

	if !edit {
		passwordHash, err := auth.Hash(u.Password)

		if err != nil {
			return err
		}

		u.Password = string(passwordHash)
	}
	return nil
}
