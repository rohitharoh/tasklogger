package cache

import (
	"github.com/tb/task-logger/backend/golang/taskapp/models"
)

type TaskCache interface {
	Set(key string, value *models.Task)
	Get(key string) *models.Task
	List (key string) []models.Task
	PSubPub (key string)
	Exists(key string) bool
	Del(key string)
	Flush(key string) error

}
