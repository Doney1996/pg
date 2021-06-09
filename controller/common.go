package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	account_repo "person/manage/repo/account"
	transLogRepo "person/manage/repo/trans_log"
	userRepo "person/manage/repo/user"
	"person/manage/utils"
	"strconv"
)

func GetUserById(c *gin.Context) {
	param := c.Query("uid")
	uId, err := strconv.ParseInt(param, 10, 64)
	utils.DealWrapErr(utils.NewErr(400, "参数错误",err))
	user := userRepo.GetById(uId)
	c.JSON(http.StatusOK, user)
}

func GetTransLogByUId(c *gin.Context){
	param := c.Query("uid")
	uId, err := strconv.ParseInt(param, 10, 64)
	utils.DealWrapErr(utils.NewErr(400, "参数错误",err))
	list := transLogRepo.GetByUId(uId)
	c.JSON(http.StatusOK, list)
}

func GetAccountByUId(c *gin.Context){
	param := c.Query("uid")
	uId, err := strconv.ParseInt(param, 10, 64)
	utils.DealWrapErr(utils.NewErr(400, "参数错误",err))
	list := account_repo.GetById(uId)
	c.JSON(http.StatusOK, list)
}
