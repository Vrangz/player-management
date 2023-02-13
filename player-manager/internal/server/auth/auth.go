package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

const validToken = "Bearer token"

type Headers struct {
	Token string `header:"Authorization"`
}

func SimpleAuthorizationMiddleware(c *gin.Context) {
	var (
		h   Headers
		err error
	)

	if err = c.ShouldBindHeader(&h); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": errors.Wrapf(err, "failed to bind header").Error()})
		return
	}

	if h.Token != validToken {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": errors.New("invalid token").Error()})
		return
	}
}
