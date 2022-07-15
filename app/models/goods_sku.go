package models

import "time"

type GoodsSku struct {
	SkuID      string    `json:"sku_id"`
	Name       string    `json:"name"`
	GoodsID    int64     `json:"goods_id"`
	Num        int64     `json:"num"`
	DeleteTime time.Time `json:"delete_time"`

	CommonTimestampField
}
