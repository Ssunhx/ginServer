package router

import (
	v1 "ginserver/api/v1"
	"ginserver/middleware"
	"ginserver/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	// 配置是否为  debug 模式
	gin.SetMode(utils.AppMode)

	// 初始化 engine ，也可以使用 New 创建，Default 会增加两个中间件
	r := gin.Default()

	// 分组路由
	// 此路由组需要验证 token
	token_router := r.Group("api/v1/")
	// 使用中间件验证 token
	token_router.Use(middleware.VerifyJwtToken())
	{
		token_router.POST("user/add", v1.AddUser)
	}

	router_v1 := r.Group("api/v1")
	// 次路由组不需要验证 token，
	{
		router_v1.POST("user/login", v1.Login)
	}

	_ = r.Run(utils.HttpPort)
}
