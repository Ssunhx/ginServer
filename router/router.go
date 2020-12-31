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
		// 上传image
		token_router.POST("image/upload", v1.UploadImg)
		// 获取当前用户所有的image
		token_router.GET("image/all", v1.GetImg)
		// 删除image
		token_router.DELETE("image/delete", v1.DeleteImg)
		// 获取已经被删除（软删除）的图片
		token_router.GET("image/deleted", v1.GetDeletedImg)
		// 删除 oss 上图片
		token_router.DELETE("image/oss/delete", v1.RealDeleteImg)
	}

	router_v1 := r.Group("api/v1/")
	// 次路由组不需要验证 token，
	{
		// 用户登录
		router_v1.POST("user/login", v1.Login)
		// 用户注册
		router_v1.POST("user/add", v1.AddUser)
	}

	_ = r.Run(utils.HttpPort)
}
