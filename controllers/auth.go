package controllers

import (
	"encoding/gob"
	"net/http"
	"strings"
	"todos/core"
	"todos/services"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/phuslu/log"
)

type AuthController struct{}

var userService = new(services.UserService)

type UserForm struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (ctrl *AuthController) Register(c *gin.Context) {
	var registration UserForm
	err := c.ShouldBindJSON(&registration)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 1, "message": "Bad request"})
		return
	}
	err = userService.CreateUser(registration.Email, registration.Password)
	if err != nil {
		log.Error().AnErr("err", err).Msg("create user failed")
		msg := "Failed to register user"
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			msg = "Email has already been registered"
		}
		c.JSON(http.StatusBadRequest, gin.H{"status": 1, "message": msg})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"status": 0})
}

func (ctrl *AuthController) Login(c *gin.Context) {
	var userForm UserForm
	err := c.ShouldBindJSON(&userForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 1, "message": "Bad request"})
		return
	}
	if user, err := userService.GetUserByEmail(userForm.Email); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": 1, "message": "Invalid email or password"})
	} else {
		if userService.CheckPassword(user, userForm.Password) {
			session := sessions.Default(c)
			gob.Register(core.User{})
			session.Set(core.UserSessionKey, user)
			if err := session.Save(); err != nil {
				log.Error().AnErr("err", err).Msg("session.Save() error")
				c.JSON(http.StatusInternalServerError, gin.H{"status": 1, "message": "Internal server error"})
			} else {
				c.JSON(http.StatusOK, gin.H{"status": 0})
			}
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"status": 1, "message": "Invalid email or password"})
		}
	}
}

func (ctrl *AuthController) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("user")
	session.Save()
	c.JSON(http.StatusOK, gin.H{"status": 0})
}

func (ctrl *AuthController) UserProfile(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(core.UserSessionKey)
	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": 1, "message": "Unauthorized"})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": 0, "user": user})
	}
}
