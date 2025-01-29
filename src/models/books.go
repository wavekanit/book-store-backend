package models

import "time"

type Books struct {
	BookID        uint      `gorm:"primaryKey;autoIncrement" json:"book_id"`
	Title         string    `gorm:"type:text;not null" json:"title"`
	Description   string    `gorm:"type:text;not null" json:"description"`
	Category      string    `gorm:"type:text;not null" json:"category"`
	PublishedDate time.Time `gorm:"type:date;not null" json:"published_date"`
	BorrowRate    int       `gorm:"default:0" json:"borrow_rate"`
	BorrowStock   int       `gorm:"default:0" json:"borrow_stock"`
	BuyPrice      int       `gorm:"default:0" json:"buy_price"`
	BuyStock      int       `gorm:"default:0" json:"buy_stock"`
}
