package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"person/manage/controller"
	"person/manage/utils"
)

func StepRouter() *gin.Engine {

	server := gin.New()
	server.Use(catch, printLog, auth)
	v1Group := server.Group("v1")
	{
		v1Group.GET("/get_UserById", controller.GetUserById)
		v1Group.GET("/get_trans_log_by_id", controller.GetTransLogByUId)
		v1Group.GET("/get_account_by_uid", controller.GetAccountByUId)
	}
	return server
}

//鉴权
func auth(context *gin.Context) {

}

//日志
func printLog(context *gin.Context) {

}

func catch(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			/*log.Printf("panic: %v\n", r)
			debug.PrintStack()*/
			c.JSON(http.StatusOK, errorToString(r))
			c.Abort()
		}
	}()
	c.Next()
}

// recover错误，转string
func errorToString(r interface{}) interface{} {
	switch v := r.(type) {
	case utils.Err:
		return map[string]interface{}{
			"code": v.Code,
			"msg":  v.Msg,
			"err":  v.Err.Error(),
		}
	case error:
		return map[string]interface{}{
			"code": 500,
			"msg":  "服务器异常",
			"err":  v.Error(),
		}
	default:
		return r.(string)
	}
}
