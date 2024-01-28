package middleware

import (
	"net/http"

	"github.com/Jhon-2801/task-manager/core/jwt"
	"github.com/gin-gonic/gin"
)

func ValidToken(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if len(token) == 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 400, "message": "Token Requerido"})
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	isValidToken, err := jwt.ProcessToken(token)
	if !isValidToken {
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 400, "message": "Token invalid", "err": err})
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		return
	}

}
