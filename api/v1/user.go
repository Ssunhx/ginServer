package v1

import (
	"ginserver/model"
	"github.com/gin-gonic/gin"
)

//var code int

func AddUser(c *gin.Context) {
	var user model.User
	_ = c.ShouldBindJSON(&user)
}
