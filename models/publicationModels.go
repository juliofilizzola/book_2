package models

import "gorm.io/gorm"

type Publication struct {
	gorm.Model
	Title       string `gorm:"not null"`
	Description string
	AuthId      string
	Content     string
	Like        int `gorm:"default:0" json:"like"`
}
