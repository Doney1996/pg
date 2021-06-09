package entity

type TransLog struct {
	Id         int64    `json:"id" db:"id"`
	AccId      int64    `json:"acc_id" db:"acc_id"`
	Amount     int64    `json:"amount" db:"amount"`
	CreateTime string `json:"create_time" db:"create_time"`
	UpdateTime string `json:"update_time" db:"update_time"`
}
