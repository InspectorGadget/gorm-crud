package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title string `binding:"required" gorm:"type:varchar(255);not null" json:"title"`
	Body  string `binding:"required" gorm:"type:varchar(255);not null" json:"body"`
}
