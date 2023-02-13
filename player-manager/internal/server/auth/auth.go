package auth

import (
	"fmt"
	"net/http"
	"player-manager/internal/server/errors"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	validUsername = "username"
	validPassword = "password"
)

func SimpleAuthorizationMiddleware(c *gin.Context) {
	u, p, ok := c.Request.BasicAuth()
	if !ok || len(strings.TrimSpace(u)) < 1 || len(strings.TrimSpace(p)) < 1 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, errors.NewError(http.StatusUnauthorized, "empty credentials", fmt.Errorf("username or password is empty")))
		return
	}

	if u != validUsername || p != validPassword {
		c.AbortWithStatusJSON(http.StatusUnauthorized, errors.NewError(http.StatusUnauthorized, "invalid credentials", fmt.Errorf("username or password is invalid")))
		return
	}
}
