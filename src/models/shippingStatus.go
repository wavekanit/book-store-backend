package models

import "time"

type ShippingStatus struct {
	OrderID       uint      `gorm:"primaryKey;not null" json:"order_id"`
	BookID        uint      `gorm:"not null" json:"book_id"`
	AddressNum    uint      `gorm:"not null" json:"address_num"`
	Status        string    `gorm:"type:text;not null" json:"status"`
	ShippingDate  time.Time `gorm:"type:timestamp;not null" json:"shipping_date"`
	DeliveredDate time.Time `gorm:"type:timestamp;not null" json:"delivered_date"`
}
