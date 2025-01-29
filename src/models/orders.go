package models

import "time"

type Order struct {
	OrderID  uint      `gorm:"primaryKey;autoIncrement" json:"order_id"`
	UserID   uint      `gorm:"not null" json:"user_id"`
	BookID   uint      `gorm:"not null" json:"book_id"`
	Type     string    `gorm:"type:text;not null" json:"type"` // e.g., 'borrow', 'buy'
	Quantity int       `gorm:"not null" json:"quantity"`
	OrderAt  time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"order_at"`
	Status   string    `gorm:"type:text;not null" json:"status"`
	DueDate  time.Time `gorm:"type:date" json:"due_date"`
}
