package v1

import (
	"ginserver/model"
	"ginserver/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

var code int

func AddUser(c *gin.Context) {
	var user model.User
	_ = c.ShouldBindJSON(&user)
	code = model.CheckUser(user.UserName)

	if code == utils.SUCCESS {
		code = model.CreateUser(&user)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    user,
		"message": utils.GetErrMsg(code),
	})

}
