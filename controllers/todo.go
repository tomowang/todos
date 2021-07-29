package controllers

import (
	"net/http"
	"todos/core"
	"todos/services"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/phuslu/log"
)

type TodosController struct{}

var todoService = new(services.TodoService)

func (ctrl *TodosController) List(c *gin.Context) {
	user := sessions.Default(c).Get(core.UserSessionKey).(core.User)
	if todos, err := todoService.List(user.ID); err == nil {
		c.JSON(http.StatusOK, gin.H{"status": 0, "data": todos})
	} else {
		log.Error().AnErr("err", err).Msg("failed to list todos for user: " + user.Email)
		c.JSON(http.StatusInternalServerError, gin.H{"status": 1, "message": "Internal server error"})
	}
}

func (ctrl *TodosController) Create(c *gin.Context) {
	todo := &core.Todo{}
	c.BindJSON(todo)
	user := sessions.Default(c).Get(core.UserSessionKey).(core.User)
	todo.UserID = user.ID
	todoService.Create(todo)
	c.JSON(http.StatusCreated, gin.H{"status": 0, "data": todo})
}

func (ctrl *TodosController) Retrieve(c *gin.Context) {}

func (ctrl *TodosController) Update(c *gin.Context) {}

func (ctrl *TodosController) Destroy(c *gin.Context) {}
