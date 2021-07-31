package main

import (
	"fmt"
	"net/http"
	"time"
	"todos/config"
	"todos/controllers"
	"todos/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	c := config.GetConfig()
	store := cookie.NewStore([]byte(c.Security.CookieSecret))
	router.Use(sessions.Sessions("session", store))

	router.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, fmt.Sprintf("%d", time.Now().Unix()))
	})
	v1 := router.Group("v1")
	{
		authCtrl := controllers.AuthController{}
		v1.POST("/register", authCtrl.Register)
		v1.POST("/login", authCtrl.Login)
		v1.POST("/logout", authCtrl.Logout)

		private := v1.Group("")
		private.Use(middleware.AuthRequired)

		private.GET("/profile", authCtrl.UserProfile)

		// todos
		todosCtrl := controllers.TodosController{}
		todos := private.Group("todos")
		todos.GET("", todosCtrl.List)
		todos.POST("", todosCtrl.Create)

		todos.GET("/:id", todosCtrl.Retrieve)
		todos.PUT("/:id", todosCtrl.Update)
		todos.DELETE("/:id", todosCtrl.Destroy)
	}
	return router
}
