package models

import "time"

type Users struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Username  string    `gorm:"type:varchar(255);not null ;unique" json:"username"`
	Password  string    `gorm:"type:text; not null" json:"password"`
	Email     string    `gorm:"type:text; not null ;unique" json:"email"`
	CreatedAt time.Time `gorm:"type:timestamp; autoCreateTime" json:"created_at"`
	Token     string    `gorm:"-" json:"token"`
}
