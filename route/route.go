package route

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"web/global"
	"web/service/basis"
)

func Route() http.Handler {
	r := gin.Default()
	r.SetTrustedProxies([]string{""})
	gin.SetMode(global.GlobalConf.Service.Env)
	//fmt.Println("username", global.GlobalConf.Mysqldb.Username)
	gin.ForceConsoleColor()
	address := fmt.Sprintf(":%d", global.GlobalConf.Service.Port)
	s := &http.Server{
		Addr:           address,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	//s.ListenAndServe()

	api := r.Group("v1")
	api.GET("/check", basis.Check)
	//r.Run()
	//s.ListenAndServe()
	err := s.ListenAndServe()
	if err != nil {
		global.GlobalLog.Error("服务启动失败，请检查端口!")
	}
	return r

}
