package middlewares

import (
	"net/http"
	"strings"

	"bionic.pro/reports-api/utils"

	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")
	if authHeader == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Authorization header missing."})
		return
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || parts[0] != "Bearer" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid Authorization header format."})
		return
	}

	token := parts[1]

	err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized: " + err.Error()})
		return
	}

	context.Next()
}
