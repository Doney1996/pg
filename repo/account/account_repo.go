package account_repo

import (
	"person/manage/database"
	"person/manage/entity"
	"person/manage/utils"
)

func GetById(id int64) *entity.Account {
	db:= database.Db.GetDb()
	var account entity.Account
	err := db.Get(&account,"select * from account where id = ?", id)
	utils.DealErr(err)
	return &account
}

func GetByUId(uId int64) *[]entity.Account {
	db:= database.Db.GetDb()
	var account  []entity.Account
	err := db.Select(&account, "select * from account where u_id = ?", uId)
	utils.DealErr(err)
	return &account
}
