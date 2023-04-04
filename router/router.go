package register

import (
	"github.com/gin-gonic/gin"
	"github.com/qqbot_zimmer/zimmer/internal/messagemgr"
)

//SetServiceRouter 设置路由
func SetServiceRouter(router *gin.Engine) {
	messagemgr.SetOnMessageRouter(router)
}
