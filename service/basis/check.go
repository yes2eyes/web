package basis

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web/global"
)

func Check(c *gin.Context) {
	//logs := utils.Zap()
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"mes":  "hahah",
	})

	global.GlobalLog.Info("检测成功")
}
