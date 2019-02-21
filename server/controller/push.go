package controller

import (
	"fmt"
	"net/http"

	"github.com/X-Sentinels/grpc-push/server/g"
	"github.com/gin-gonic/gin"
)

func Validate(c *gin.Context) {
	apiKey := c.Request.Header.Get("X-API-KEY")
	if apiKey != g.Config().Http.X_API_KEY {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"msg":     "X-API-KEY IS Wrong",
		})
		c.Abort()
		return
	}
}

func PushNotification(c *gin.Context) {
	clientName := c.Request.Header.Get("clientName")
	if !g.InArray(clientName, g.Config().AliveClients) {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"msg":     "ClientName is not in the alive list",
		})
		return
	}
	body, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"msg":     fmt.Sprintf("Get Request Body Failed %s", err.Error()),
		})
		return
	}
	notif := g.Message{ClientName: clientName, Notification: string(body)}
	select {
	case g.NotifMessage <- notif:
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"msg":     "ok",
		})
		return
	default:
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"msg":     "Too Manny Message in Channel",
		})
		return
	}
}

func GetClients(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"msg":     "ok",
		"data":    g.Config().AliveClients,
	})
	return

}
