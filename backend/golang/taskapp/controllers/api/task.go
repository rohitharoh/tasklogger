package api

import (
	"encoding/json"

	_ "fmt"
	log "github.com/sirupsen/logrus"
	"github.com/tb/task-logger/backend/golang/common-packages/system"
	cache "github.com/tb/task-logger/backend/golang/taskapp/cache"
	"github.com/tb/task-logger/backend/golang/taskapp/models"
	"github.com/tb/task-logger/backend/golang/taskapp/services"
	_"github.com/tb/task-logger/backend/golang/taskapp/validations"
	"github.com/zenazn/goji/web"
	"net/http"
	"strconv"
)

type Controller struct {
	TaskController
}

var (
	taskService services.TaskService
	taskCache   cache.TaskCache
)

type TaskController interface {
	TaskDetails(c web.C, w http.ResponseWriter, r *http.Request, logger *log.Entry) ([]byte, error)
	ListTask(c web.C, w http.ResponseWriter, r *http.Request, logger *log.Entry) ([]byte, error)
	AddTask(c web.C, w http.ResponseWriter, r *http.Request, logger *log.Entry) ([]byte, error)
}

func NewPostController(service services.TaskService, cache cache.TaskCache) TaskController {
	taskService = service
	taskCache = cache
	return &Controller{}
}

func (controller *Controller) AddTask( c web.C, w http.ResponseWriter, r *http.Request, logger *log.Entry) ([]byte, error) {
	decoder := json.NewDecoder(r.Body)
	var addTaskInput models.AddTaskInput
	err := decoder.Decode(&addTaskInput)
	if err != nil {
		logger.Error(err)
		return nil, system.InvalidPayloadError
	}
	logger.Info(addTaskInput)
	response, err := services.AddTask(logger, addTaskInput, c.Env["emailId"].(string))
	return response, err
}

func (controller *Controller) ListTask( c web.C, w http.ResponseWriter, r *http.Request, logger *log.Entry) ([]byte, error) {
	decoder := json.NewDecoder(r.Body)
	var listTaskParam map[string]string
	err := decoder.Decode(&listTaskParam)
	if err != nil {
		logger.Error(err)
		return nil, system.InvalidPayloadError
	}

	taskStatue, keyExists := listTaskParam["status"]
	if !keyExists {
		return nil, system.InvalidPayloadError
	}
	if taskStatue == "" {
		taskStatue = system.TASK_STATUS_PENDING
	}

	skip := r.URL.Query().Get("skip")
	skipValue, _ := strconv.Atoi(skip)

	response, err := services.ListTask(logger, taskStatue, c.Env["emailId"].(string), skipValue)
	return response, err
}

func (controller *Controller) TaskDetails( c web.C, w http.ResponseWriter, r *http.Request, logger *log.Entry) ([]byte, error) {
	decoder := json.NewDecoder(r.Body)
	var taskParam map[string]string
	err := decoder.Decode(&taskParam)
	if err != nil {
		logger.Error(err)
		return nil, system.InvalidPayloadError
	}

	recordId, keyExists := taskParam["recordId"]
	if !keyExists {
		return nil, system.InvalidPayloadError
	}



	response, err := services.TaskDetails(logger, c.Env["emailId"].(string), recordId)
	return response, err
}

