package middleware

import (
	"net/http"
	"todos/core"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// AuthRequired is a simple middleware to check the session
func AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(core.UserSessionKey)
	if user == nil {
		// Abort the request with the appropriate error code
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": 1, "message": "Unauthorized"})
		return
	}
	// Continue down the chain to handler etc
	c.Next()
}
