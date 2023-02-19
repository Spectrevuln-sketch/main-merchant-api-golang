package models

type Customer struct {
	Id     int64  `gorm:"primaryKey" json:"id"`
	Name   string `gorm:"type:varchar(300)" json:"nama_merchant"`
	Alamat string `gorm:"type:varchar(300)" json:"alamat"`
	Phone  string `gorm:"type:varchar(300)" json:"phone"`
}
