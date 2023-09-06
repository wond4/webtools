package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wond4/webtools/router"
)

func main() {
	//1.创建路由
	r := gin.Default()
	//2.绑定路由规则，执行的函数
	router.Setup(r)
	//3.监听端口，默认8080
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
