package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"messagesystem/internal/storage"
	"net"
	"net/http"
)

type API struct {
	Storage *storage.Storage
	TCPAddr string
}

func (a *API) Routes() *gin.Engine {
	r := gin.Default()

	r.GET("/messages", func(c *gin.Context) {
		msgs, err := a.Storage.GetMessages()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, msgs)
	})

	r.POST("/send", func(c *gin.Context) {
		var body struct {
			Message string `json:"message"`
		}
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		conn, err := net.Dial("tcp", a.TCPAddr)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		defer conn.Close()

		fmt.Fprintf(conn, body.Message+"\n")
		c.JSON(200, gin.H{"status": "sent"})
	})

	return r
}
