package ip

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ClientIp(context *gin.Context) {
	context.String(http.StatusOK, context.ClientIP())
}
