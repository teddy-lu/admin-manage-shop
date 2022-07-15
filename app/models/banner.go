package models

type Banner struct {
	BannerID   int64  `json:"banner_id"`
	Name       string `json:"name"`
	StoreID    int64  `json:"store_id"`
	Img        string `json:"img"`
	Sort       int64  `json:"sort"`
	Background string `json:"background"`

	CommonTimestampField
}
