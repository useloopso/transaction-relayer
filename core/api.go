package core

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/useloopso/transaction-relayer/config"
	"github.com/useloopso/transaction-relayer/types"
)

func RunRouter() {
	r := configRouter()
	r.Run(fmt.Sprintf(":%d", config.Get().Port))
}

func configRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", ping)
	r.POST("/execute", execute)
	return r
}

func execute(c *gin.Context) {
	var payload types.ExecutePayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": err.Error()})
		return
	}

	if err := ExecuteRelayCall(payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": err.Error()})
		return
	}
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, "pong")
}
