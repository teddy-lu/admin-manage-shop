package models

type User struct {
	Password string `json:"-"`

	CommonTimestampField
}
