package models

type Address struct {
	UserID      uint   `gorm:"primaryKey" json:"user_id"`
	AddressNum  uint   `gorm:"primaryKey" json:"address_num"`
	AddressName string `gorm:"type:text;not null" json:"address_name"`
	Detail      string `gorm:"type:text;not null" json:"detail"`
	City        string `gorm:"type:text;not null" json:"city"`
	Province    string `gorm:"type:text;not null" json:"province"`
	ZipCode     string `gorm:"type:text;not null" json:"zip_code"`
	Country     string `gorm:"type:text;not null" json:"country"`
}
