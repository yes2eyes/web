package route

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"web/service/basis"
)

func Route() http.Handler {
	r := gin.Default()
	r.SetTrustedProxies([]string{""})
	gin.ForceConsoleColor()
	s := &http.Server{
		Addr:           ":8081",
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
	s.ListenAndServe()
	return r

}
