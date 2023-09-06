package router

import (
	"github.com/gin-gonic/gin"
	"github.com/wond4/webtools/app/ip"
)

func Setup(engine *gin.Engine) {
	// ip
	engine.Any("/ip", ip.ClientIp)
}
