package models

type GoodsImage struct {
	ImageID int64  `json:"image_id"`
	Src     string `json:"src"`
	GoodsID int64  `json:"goods_id"`
	Sort    int64  `json:"sort"`

	CommonTimestampField
}
