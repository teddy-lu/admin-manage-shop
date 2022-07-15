package models

type OrderDetail struct {
	DetailID int64  `json:"detail_id"`
	OrderID  int64  `json:"order_id"`
	SkuID    string `json:"sku_id"`
	CateID   int64  `json:"cate_id"`
	GoodsID  int64  `json:"goods_id"`
	Img      string `json:"img"`
	Num      int64  `json:"num"`

	CommonTimestampField
}
