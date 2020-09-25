package v1

import (
	"ginserver/middleware"
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

func Login(c *gin.Context) {
	// 登录接口，
	var data model.User
	_ = c.ShouldBindJSON(&data)

	var token string
	var code int

	// 检查用户名和密码
	code = model.CheckLogin(data.UserName, data.Password)
	if code == utils.SUCCESS {
		// 用户校验通过之后，生成 token，若验证未通过， token 为空
		token, code = middleware.GenJwtToken(data.UserName)
	} else {
		code = utils.ERROR
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": utils.GetErrMsg(code),
		"data":    token,
	})
}
