package models

import "time"

type BaseModel struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement;" json:"id,omitempty"`
}

type CommonTimestampField struct {
	CreateTime time.Time `gorm:"column:create_time;index" json:"create_time,omitempty"`
	UpdateTime time.Time `gorm:"column:update_time;index" json:"update_time,omitempty"`
}
