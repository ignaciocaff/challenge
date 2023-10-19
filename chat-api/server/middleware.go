package server

import (
	"chatjobsity/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthRequired(c *gin.Context) {
	sessionManager := c.MustGet("session_manager").(*utils.SessionManager)

	currentPath := c.FullPath()
	if currentPath == "/api/auth" || currentPath == "/api/auth/me" {
		c.Next()
		return
	}
	cookie, err := c.Cookie("jobsity")
	if err != nil {
		if err == http.ErrNoCookie {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		}
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}
	sessions := sessionManager.GetSessionsMap()
	userSession, exists := sessions.Load(cookie)
	session := userSession.(utils.Session)

	if !exists {
		// If the session token is not present in session map, return an unauthorized error
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}
	// If the session is present, but is invalid, we can delete the session, and return an unauthorized status
	if !session.IsValid {
		sessions.Delete(cookie)
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Session invalid"})
		return
	}
	c.Next()
}
