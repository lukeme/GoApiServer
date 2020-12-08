package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wuyan94zl/api/app/controllers/admin"
)

// 注册路由列表
func ApiRouter(router *gin.RouterGroup) {
	router.POST("/admin/login", admin.Login) // 登录
}
