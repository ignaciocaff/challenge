package handlers

import (
	"chatjobsity/models"
	"chatjobsity/services"
	"chatjobsity/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AuthHandler struct {
	authService    services.AuthService
	sessionManager *utils.SessionManager
}

func New(authService services.AuthService, sessionManager *utils.SessionManager) *AuthHandler {
	return &AuthHandler{authService: authService, sessionManager: sessionManager}
}

func (h *AuthHandler) Login(c *gin.Context) {
	cookie, _ := c.Cookie("jobsity")

	var login LoginJson

	if err := c.ShouldBind(&login); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error getting login data"})
		return
	}

	user, err := h.authService.Login(login.Username, login.Password)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	sessionToken := uuid.NewString()

	if cookie == "" {
		h.sessionManager.AddSession(sessionToken, utils.Session{Id: user.Id, Username: user.Username, IsValid: true})
		c.SetCookie("jobsity", sessionToken, 3600, "/", "", true, true)
		c.JSON(http.StatusOK, user)
		return
	}

	session, err := h.sessionManager.GetSession(cookie)

	if err != nil {
		// If the session token is not present in session map, return a new session
		h.sessionManager.AddSession(sessionToken, utils.Session{Id: user.Id, Username: user.Username, IsValid: true})
		c.SetCookie("jobsity", sessionToken, 3600, "/", "", true, true)
		c.JSON(http.StatusOK, models.User{Id: user.Id, Username: user.Username})
		return
	} else {
		if !session.IsValid {
			// If the session is present, but is invalid, we can delete the session, and return a new one
			h.sessionManager.RemoveSession(cookie)
			h.sessionManager.AddSession(sessionToken, utils.Session{Id: user.Id, Username: user.Username, IsValid: true})
		} else {
			// If the session is present, and is valid, we can return the same session
			sessionToken = cookie
		}
		c.SetCookie("jobsity", sessionToken, 3600, "/", "", true, true)
		c.JSON(http.StatusOK, models.User{Id: user.Id, Username: user.Username})
		return
	}
}

func (h *AuthHandler) Me(c *gin.Context) {
	cookie, _ := c.Cookie("jobsity")

	if cookie == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	session, _ := h.sessionManager.GetSession(cookie)

	if session != nil && session.IsValid {
		c.JSON(http.StatusOK, models.User{Id: session.Id, Username: session.Username})
		return
	}
	c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
}

func (h *AuthHandler) Logout(c *gin.Context) {
	cookie, err := c.Cookie("jobsity")
	if err != nil {
		if err == http.ErrNoCookie {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		}
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}
	h.sessionManager.InvalidSession(cookie)

	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}

func (h *AuthHandler) SignUp(c *gin.Context) {
	var user models.User

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error getting user data"})
		return
	}

	if err := h.authService.SignUp(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

type LoginJson struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
