package models

import "time"

type Review struct {
	ReviewID    uint      `gorm:"primaryKey;autoIncrement" json:"review_id"`
	UserID      uint      `gorm:"not null" json:"user_id"`
	BookID      uint      `gorm:"not null" json:"book_id"`
	Description string    `gorm:"type:text;not null" json:"description"`
	ReviewDate  time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"review_date"`
	Rating      int       `gorm:"not null" json:"rating"`
}
