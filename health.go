package zflag

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Health
// @Description: 服务健康检查，暂时为空动作
// @param c
func Health(c *gin.Context) {
	c.JSON(http.StatusOK, Response(1, "", "done", "", ""))
	c.Abort()
}
