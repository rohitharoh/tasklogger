package services

import (
	"encoding/json"
	"fmt"
	"github.com/pborman/uuid"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/tb/task-logger/backend/golang/common-packages/system"
	cache "github.com/tb/task-logger/backend/golang/taskapp/cache"
	"github.com/tb/task-logger/backend/golang/taskapp/models"
	"github.com/tb/task-logger/backend/golang/taskapp/validations"

	"gopkg.in/mgo.v2/bson"
	"time"
)
type TaskService interface {
	AddTask(logger *logrus.Entry, createTaskInput models.AddTaskInput, emailId string) ([]byte, error)
	ListTask(logger *logrus.Entry, taskStatus string, emailId string) ([]byte, error)
	TaskDetails(logger *logrus.Entry, emailId string, recordId string) ([]byte, error)
}

func AddTask(logger *logrus.Entry, createTaskInput models.AddTaskInput, emailId string) ([]byte, error) {

	isValid := Validationspackage.ValidateEmail(emailId)
	if !isValid {
		return nil, system.InvalidEmailErr
	}

	err := Validationspackage.ValidateAddTaskInput(logger, createTaskInput)
	if err != nil {
		return nil, err
	}

	taskObj := models.Task{
		Id:          uuid.New(),
		Title:       createTaskInput.Title,
		ScheduledOn: createTaskInput.ScheduledOn,
		Description: createTaskInput.Description,
		EmailId:     emailId,
		Status:      system.TASK_STATUS_PENDING,
		CreatedOn:   time.Now(),
		ModifiedOn:  time.Now(),
	}

	fmt.Println("_id", taskObj.Id)

//cache.NewRedisCache("127.0.0.1", 0, system.REDIS_DEFAULT_EXPIRATION_TIME).Flush("")
	client := cache.NewRedisCache("127.0.0.1", 0, system.REDIS_DEFAULT_EXPIRATION_TIME)

	client.Set(system.TASKS_COLLECTION + ":" + taskObj.Id, &taskObj)

	collectionName := system.TASKS_COLLECTION
	databaseName := system.GetDatabaseName(collectionName)
	sessionDb := system.TbAppContext.MongoDBSessionMap[databaseName].Clone()
	defer sessionDb.Close()
	collection := sessionDb.DB(databaseName).C(collectionName)
	err = collection.Insert(&taskObj)
	if err != nil {
		return nil, err
	}
	response := make(map[string]interface{})
	response["message"] = "Task created successfully"
	finalResponse, _ := json.Marshal(response)
	return finalResponse, nil

}


func ListTask(logger *logrus.Entry, taskStatus string, emailId string, skip int) ([]byte, error) {

	totalskip := skip
	skip = skip * system.LIMIT_RECORD_FOR_INVITES
	isValid := Validationspackage.ValidateEmail(emailId)
	if !isValid {
		return nil, system.InvalidEmailErr
	}

	validTaskStatus, _ := system.Contains(taskStatus, [...]string{system.TASK_STATUS_PENDING, system.TASK_STATUS_DONE, system.TASK_STATUS_DELETED, system.TASK_STATUS_ALL})
	if !validTaskStatus {
		return nil, system.InvalidStatusErr
	}

	var taskList []models.Task

	key := system.TASKS_COLLECTION + ":" + "*"
	taskList = cache.NewRedisCache(viper.GetString("redis.addr"), 0, system.REDIS_DEFAULT_EXPIRATION_TIME).List(key)
	cache.NewRedisCache(viper.GetString("redis.addr"), 0, system.REDIS_DEFAULT_EXPIRATION_TIME).PSubPub(key)

	var totalCount int
	if len(taskList) == 0 {
		fmt.Println("getting from db collection")
		collectionName := system.TASKS_COLLECTION
		databaseName := system.GetDatabaseName(collectionName)
		sessionDb := system.TbAppContext.MongoDBSessionMap[databaseName].Clone()
		defer sessionDb.Close()
		collection := sessionDb.DB(databaseName).C(collectionName)

		queryCondition := bson.M{"emailId": emailId}
		if taskStatus != "all" {
			queryCondition["status"] = taskStatus
		}

		err := collection.Find(queryCondition).All(&taskList)
		if err != nil {
			return nil, err
		}

		totalCount, err = collection.Find(queryCondition).Count()
		if err != nil {
			return nil, err
		}

	}

	var isMoreList bool
	if totalCount > skip+system.LIMIT_RECORD_FOR_INVITES {
		isMoreList = true
	} else {
		isMoreList = false
	}
	taskListResp := make([]map[string]string, 0)
	//var taskListResp []map[string]string
	if len(taskList) > 0 {

		count := 0
		limit := system.LIMIT_RECORD_FOR_INVITES
		s := totalskip * limit
		for n, taskObj := range taskList {

			if n >= s {
				if count < limit {

					createdOnDate := taskObj.CreatedOn.Format("2006-01-02")
					completedOnDate := ""
					if taskObj.Status != system.TASK_STATUS_PENDING && taskObj.CompletedOn.Format("2006-01-02") != "0001-01-01" {
						completedOnDate = taskObj.CompletedOn.Format("2006-01-02")
					}
					taskListResp = append(taskListResp, map[string]string{
						"id":              taskObj.Id,
						"title":           taskObj.Title,
						"scheduledOn":     taskObj.ScheduledOn,
						"description":     taskObj.Description,
						"emailId":         taskObj.EmailId,
						"status":          taskObj.Status,
						"createdOnDate":   createdOnDate,
						"completedOnDate": completedOnDate,
					})
				}
				count++
			}
		}
	}
	response := make(map[string]interface{})
	response["isMoreList"] = isMoreList
	response["task_list"] = taskListResp

	finalResponse, _ := json.Marshal(response)
	return finalResponse, nil
}

func TaskDetails(logger *logrus.Entry, emailId string, recordId string) ([]byte, error) {

	isValid := Validationspackage.ValidateEmail(emailId)
	if !isValid {
		return nil, system.InvalidEmailErr
	}
	if recordId == "" {
		return nil, system.NoRecordIdErr
	}
	var taskDetails *models.Task
	taskDetails = cache.NewRedisCache("127.0.0.1:6379", 0, system.REDIS_DEFAULT_EXPIRATION_TIME).Get(recordId)
	fmt.Println("post-----cache------->", taskDetails)

	key := system.TASKS_COLLECTION + ":" + recordId
   cache.NewRedisCache(viper.GetString("redis.addr"), 0, system.REDIS_DEFAULT_EXPIRATION_TIME).PSubPub(key)
	if taskDetails == nil {
		fmt.Println("post is nil")
		collectionName := system.TASKS_COLLECTION
		databaseName := system.GetDatabaseName(collectionName)
		sessionDb := system.TbAppContext.MongoDBSessionMap[databaseName].Clone()
		defer sessionDb.Close()
		collection := sessionDb.DB(databaseName).C(collectionName)

		err := collection.Find(bson.M{"emailId": emailId, "_id": recordId}).One(&taskDetails)
		if err != nil {
			logger.Error(err)
			if err.Error() == system.NotFoundErr.Error() {
				return nil, system.InvalidRecordId
			} else {
				return nil, err
			}
		}

		cache.NewRedisCache(viper.GetString("redis.addr"), 0, system.REDIS_DEFAULT_EXPIRATION_TIME).Set(taskDetails.Id, taskDetails)
	}
	completedOnDate := ""
	if taskDetails.Status != system.TASK_STATUS_PENDING && taskDetails.CompletedOn.Format("2006-01-02") != "0001-01-01" {
		completedOnDate = taskDetails.CompletedOn.Format("2006-01-02")
	}
	response := make(map[string]interface{})
	response["task_detail"] = map[string]string{
		"id":              taskDetails.Id,
		"title":           taskDetails.Title,
		"scheduledOn":     taskDetails.ScheduledOn,
		"description":     taskDetails.Description,
		"emailId":         taskDetails.EmailId,
		"status":          taskDetails.Status,
		"createdOnDate":   taskDetails.CreatedOn.Format("2006-01-02"),
		"completedOnDate": completedOnDate,
	}
	finalResponse, _ := json.Marshal(response)
	return finalResponse, nil
}
