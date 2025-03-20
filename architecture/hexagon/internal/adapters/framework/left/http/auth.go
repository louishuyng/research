package http

import (
	contracts "coffee/internal/adapters/framework/left/http/contracts"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a Adapter) registerAuthRoutes() {
	a.r.POST("/signup", func(c *gin.Context) {
		var req contracts.SignUpRequest

		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		authData, err := a.AuthPort.SignUp(c.Request.Context(), req.Email, req.Password)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		resp := contracts.SignUpResponse{
			User: contracts.AuthUser{
				Email: authData.User.Email,
			},
			AccessToken:  authData.AccessToken,
			RefreshToken: authData.RefreshToken,
		}

		c.JSON(http.StatusOK, resp)
	})

	a.r.POST("/login", func(c *gin.Context) {
		var req contracts.SignInRequest

		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		authData, err := a.AuthPort.SignIn(c.Request.Context(), req.Email, req.Password)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		resp := contracts.SignInResponse{
			User: contracts.AuthUser{
				Email: authData.User.Email,
			},
			AccessToken:  authData.AccessToken,
			RefreshToken: authData.RefreshToken,
		}

		c.JSON(http.StatusOK, resp)
	})
}
