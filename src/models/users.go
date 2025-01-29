package models

import "time"

type Users struct {
	UserID     uint      `gorm:"primaryKey;autoIncrement" json:"user_id"`
	F_name     string    `gorm:"type:text;not null" json:"f_name"`
	L_name     string    `gorm:"type:text;not null" json:"l_name"`
	Tel        string    `gorm:"type:text;not null" json:"tel"`
	Email      string    `gorm:"type:text;not null;unique" json:"email"`
	Username   string    `gorm:"type:text;not null;unique" json:"username"`
	Password   string    `gorm:"type:text;not null" json:"password"`
	CreatedAt  time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	Token      string    `gorm:"-" json:"token"`
	Auth_level uint      `gorm:"default:0" json:"auth_level"`
}
