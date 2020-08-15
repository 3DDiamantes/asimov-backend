package middleware

import (
	"asimov-backend/internal/jwt"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware interface {
	Check(c *gin.Context)
}

type authMiddleware struct {
}

func NewAuthMiddleware() AuthMiddleware {
	return &authMiddleware{}
}

func (a *authMiddleware) Check(c *gin.Context) {
	bearerToken := c.GetHeader("Authorization")

	bearerTokenSplit := strings.Split(bearerToken, " ")

	if len(bearerTokenSplit) != 2 || bearerTokenSplit[0] != "Bearer" || len(strings.Split(bearerTokenSplit[1], ".")) != 3 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid Authorization Code"})
	}

	if !jwt.Verify(bearerTokenSplit[1]) {
		c.AbortWithError(http.StatusNotFound, errors.New("Signature doesn't match for token "+bearerTokenSplit[1]))
	}

	c.Next()
}
