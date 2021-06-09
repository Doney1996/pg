package entity

type Account struct {
	Id         int64     `json:"id" db:"id"`
	UId        int64     `json:"u_id" db:"u_id"`
	Name       string    `json:"name" db:"name"`
	AcType     int64      `json:"ac_type" db:"ac_type"`
	Balance    int64     `json:"balance" db:"balance"`
	CreateTime string `json:"create_time" db:"create_time"`
	UpdateTime string `json:"update_time" db:"update_time"`
}
