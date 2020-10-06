package system

import "time"

const (
	USER_CONTEXT              = "userContext"
	EMPTY_STRING              = ""
	StatusUnprocessableEntity = 422

	INSERT_COLLECTION             = "insert"
	TASKS_COLLECTION              = "task"
	LIMIT_RECORD_FOR_INVITES      = 12
	REDIS_DEFAULT_EXPIRATION_TIME = 5 * time.Minute
	TASK_STATUS_PENDING           = "pending"
	TASK_STATUS_DONE              = "done"
	TASK_STATUS_DELETED           = "deleted"
	TASK_STATUS_ALL               = "all"
	CUSTOM_DATE_FORMAT            = "02 Jan, 2006"
)
