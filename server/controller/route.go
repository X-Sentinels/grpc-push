package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func StartGin(listen string) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(Validate)

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ds-api version 0.1")
	})
	r.POST("/api/v1/push", PushNotification)
	r.GET("/api/v1/clients", GetClients)
	r.Run(listen)
}
