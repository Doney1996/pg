package entity

type User struct {
	Id         int64    `json:"id" db:"id"`
	Username   string `json:"username" db:"username"`
	Password   string `json:"password" db:"password"`
	CreateTime string `json:"create_time" db:"create_time"`
	UpdateTime string `json:"update_time" db:"update_time"`
}



