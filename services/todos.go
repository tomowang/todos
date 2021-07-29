package services

import (
	"todos/core"
	database "todos/db"

	"github.com/phuslu/log"
)

type TodoService struct{}

func (s *TodoService) List(userID uint) (*[]core.Todo, error) {
	log.Debug().Msg("todo.List")
	db := database.GetDB()
	todos := []core.Todo{}
	if err := db.Where("user_id = ?", userID).Find(&todos).Error; err != nil {
		return nil, err
	} else {
		return &todos, nil
	}
}

func (s *TodoService) Create(todo *core.Todo) error {
	log.Debug().Msg("todo.Create")
	db := database.GetDB()
	err := db.Create(todo).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *TodoService) Retrieve(id uint) (*core.Todo, error) {
	log.Debug().Uint("id", id).Msg("todo.Retrieve")
	db := database.GetDB()
	todo := core.Todo{}
	if err := db.First(&todo, id).Error; err != nil {
		return nil, err
	} else {
		return &todo, nil
	}
}
