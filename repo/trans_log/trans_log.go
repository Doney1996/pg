package trans_log_repo

import (
	"person/manage/database"
	"person/manage/entity"
	"person/manage/utils"
)

// AddTransLogs 批量增加多条交易
func AddTransLogs(logList []entity.TransLog) {
	db, err := database.Db.GetDb().Beginx()
	utils.DealErr(err)

	var sql = "INSERT INTO trans_log (acc_id, amount, create_time, update_time) VALUES (?, ?, ?, ?);"
	for _, log := range logList {
		_, err := db.Exec(sql, log.AccId, log.Amount, log.CreateTime, log.UpdateTime)
		utils.DealErrAndRollback(db, err)
	}
}

// AddTransLog 增加单条交易
func AddTransLog(log entity.TransLog) {
	var logList []entity.TransLog
	logList = append(logList, log)
	AddTransLogs(logList)
}

// UpdateTransLog 根据id修改交易
func UpdateTransLog(log entity.TransLog) {
	var sql = "UPDATE trans_log SET acc_id = ?, amount = ?, create_time = ?, update_time = ? WHERE id = 1;"
	db := database.Db.GetDb()
	_, err := db.Exec(sql, log.AccId, log.Amount, log.CreateTime, log.UpdateTime, log.Id)
	utils.DealErr(err)
}

func GetByUId(uId int64) *[]entity.TransLog {
	var sql  = "select * from trans_log where uid = ?"
	db := database.Db.GetDb()
	var list []entity.TransLog
	err := db.Select(&list, sql, uId)
	utils.DealWrapErr(utils.NewErr(500,"保存数据库失败",err))
	return &list
}
