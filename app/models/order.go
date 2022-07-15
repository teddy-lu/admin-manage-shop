package models

import "time"

type Order struct {
	OrderID     int64     `json:"order_id"`
	OrderSn     string    `json:"order_sn"`
	PayPrice    float64   `json:"pay_price"`
	UserID      int64     `json:"user_id"`
	GoodsNum    int64     `json:"goods_num"`
	AddressID   int64     `json:"address_id"`
	AddressJson string    `json:"address_json"`
	DeleteTime  time.Time `json:"delete_time"`
	Status      int64     `json:"status"`
	PayTime     time.Time `json:"pay_time"`

	CommonTimestampField
}
