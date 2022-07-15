package models

import "time"

type Category struct {
	CateID     int64     `json:"cate_id"`
	Name       string    `json:"name"`
	StoreID    int64     `json:"store_id"`
	Pid        int64     `json:"pid"`
	DeleteTime time.Time `json:"delete_time"`

	CommonTimestampField
}
