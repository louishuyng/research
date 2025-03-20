package http

import (
	"log"
	"net/http"

	ports "coffee/internal/ports/left/http"

	"github.com/gin-gonic/gin"
)

type ApiPorts struct {
	AuthPort ports.AuthAPIPort
}

type Adapter struct {
	r *gin.Engine
	ApiPorts
}

func NewAdapter(apiPorts ApiPorts) Adapter {
	return Adapter{
		r:        gin.Default(),
		ApiPorts: apiPorts,
	}
}

func (a Adapter) Run(port string) {
	a.r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	a.registerAuthRoutes()

	err := a.r.Run(":" + port)

	if err != nil {
		log.Fatalf("failed to run server: %v", err)
	}

	log.Printf("server running on port %s", port)
}
