package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sixi-store/users/app"
)

// R gin启动存储变量
var R *gin.Engine

func init() {
	R = gin.Default()
}

//APP app路由
func APP() {
	R.GET("/user", app.UserList)
	R.GET("/user/:name", app.UserInfo)
}
