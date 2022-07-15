package models

type Cart struct {
	CartID  int64  `json:"cart_id"`
	SkuID   string `json:"sku_id"`
	CateID  int64  `json:"cate_id"`
	UserID  int64  `json:"user_id"`
	GoodsID int64  `json:"goods_id"`
	Num     int64  `json:"num"`

	CommonTimestampField
}
