package models

import "time"

type Address struct {
	AddressID    int64     `json:"address_id" gorm:"column:address_id;primaryKey;autoIncrement;"`
	UserID       int64     `json:"user_id" gorm:"column:user_id"`
	Address      string    `json:"address" gorm:"column:address"`
	City         string    `json:"city" gorm:"column:city"`
	Country      string    `json:"country" gorm:"column:country"`
	PostCode     int64     `json:"post_code" gorm:"column:post_code"`
	Area         string    `json:"area" gorm:"column:area"`
	IsDefault    int64     `json:"is_default" gorm:"column:is_default"`
	ContactPhone string    `json:"contact_phone" gorm:"column:contact_phone"`
	ContactName  string    `json:"contact_name" gorm:"column:contact_name"`
	DeleteTime   time.Time `json:"delete_time" gorm:"column:delete_time"`

	CommonTimestampField
}
