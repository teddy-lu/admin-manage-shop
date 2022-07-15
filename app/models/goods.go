package models

import "time"

type Goods struct {
	GoodsID    int64     `json:"goods_id"`
	Name       string    `json:"name"`
	CateID     int64     `json:"cate_id"`
	SalesNum   int64     `json:"sales_num"`
	SalesPrice float64   `json:"sales_price"`
	ShowPrice  float64   `json:"show_price"`
	Unit       string    `json:"unit"`
	Src        string    `json:"src"`
	DeleteTime time.Time `json:"delete_time"`

	CommonTimestampField
}
