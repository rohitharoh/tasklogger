package models

import "time"



type AddTaskInput struct {
	Title 		  string   			`json:"title" bson:"title"`
	ScheduledOn 	  string			`json:"scheduledOn" bson:"scheduledOn"`
	Description       string			`json:"description" bson:"description"`

}

type Task struct {
	Id		  string       			`json:"_id" bson:"_id" `
	Title 		  string   			`json:"title" bson:"title"`
	ScheduledOn 	  string			`json:"scheduledOn" bson:"scheduledOn"`
	Description       string			`json:"description" bson:"description"`
	EmailId 	  string			`json:"emailId" bson:"emailId"`
	Status            string                        `json:"status" bson:"status"`
	CreatedOn         time.Time                     `json:"createdOn" bson:"createdOn"`
	CompletedOn       time.Time                     `json:"completedOn" bson:"completedOn"`
	ModifiedOn        time.Time                     `json:"modifiedOn" bson:"modifiedOn"`
}

