package models

import (
	"time"
)

// Article represents the articles table in the database
type Article struct {
	ID      uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Author  string    `gorm:"type:text;not null" json:"author"`
	Title   string    `gorm:"type:text;not null" json:"title"`
	Body    string    `gorm:"type:text;not null" json:"body"`
	Created time.Time `gorm:"autoCreateTime" json:"created"`
}
