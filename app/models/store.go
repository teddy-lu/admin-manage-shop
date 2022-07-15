package models

import "time"

type Store struct {
	StoreID    int64     `json:"store_id"`
	Name       string    `json:"name"`
	Address    string    `json:"address"`
	City       string    `json:"city"`
	Province   string    `json:"province"`
	PostCode   int64     `json:"post_code"`
	DeleteTime time.Time `json:"delete_time"`

	CommonTimestampField
}
