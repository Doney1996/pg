package user_repo

import (
	"person/manage/database"
	"person/manage/entity"
	"person/manage/utils"
)


func GetById(id int64) *entity.User {
	db := database.Db.GetDb()
	var user entity.User
	err := db.Get(&user,"select * from user where id = ?", id)
	utils.DealErr(err)
	return &user
}
