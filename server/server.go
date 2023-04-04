package server

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/qqbot_zimmer/zimmer/internal/config/initparams"
	register "github.com/qqbot_zimmer/zimmer/router"
)

//StartHTTPServer 开启http服务
func StartHTTPServer() {
	port := os.Getenv("PORT")
	if port == "" {
		port = initparams.DefaultPort
	}

	//参数初始化
	initparams.GetInstance()

	router := InitGinEngine()

	runHTTPServer(port, router)
}

//InitGinEngine 初始化Gin
func InitGinEngine() *gin.Engine {
	//生成gin
	router := gin.Default()

	//注册路由
	register.SetServiceRouter(router)
	return router
}

//runHttpServer 启动gin引擎运行
func runHTTPServer(port string, router *gin.Engine) {
	err := router.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		panic(fmt.Sprintf("Failed to start server: %v", err))
	}

}
