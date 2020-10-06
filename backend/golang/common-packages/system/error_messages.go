package system

import "errors"

var (
	InternalServerError = errors.New("Sorry, Something went wrong.")
	RedisNewUserLoginError = errors.New("redis: nil")
	MapKeyError = errors.New("MapKeyError")
	InvalidPayloadError = errors.New("Please check your request. You are missing some required details.")

	//Auth related messages
	UnauthorisedErr = errors.New("Unauthorized request")
	EmailMissingErr = errors.New("Email Id missing")
	NotFoundErr = errors.New("not found") // do not change this error message text
	InvalidRecordId = errors.New("Invalid record reference provided")

	InvalidEmailErr = errors.New("Enter valid email id")
	InvalidTitleErr = errors.New("Enter valid task title")
	InvalidScheduledOnErr = errors.New("Provide valid date")
	InvalidDateTimeFormatErr = errors.New("Provide valid date & time in proper format")
	InvalidDescriptionErr = errors.New("Description cannot be empty")
	NoUserFoundErr = errors.New("No user found")
	NoTaskFoundErr = errors.New("No records found error")
	NoListFoundErr = errors.New("No list found error")
	NoRecordIdErr = errors.New("Please provide record id")
	NoStatusErr = errors.New("Enter status")
	InvalidStatusChangeErr = errors.New("Provide valid status")
	InvalidStatusErr = errors.New("Invalid status reference provided")
	AlreadyDeletedRecordErr = errors.New("The record has already been deleted")
	NotInPendingStateErr = errors.New("Can not mark the referred task as done, since it is not in pending state.")

	InvalidDateFormatErr = errors.New("Provide valid date in proper format")
)

func GetErrorMessagesMap() (map[error]bool) {

	errorMessageMap := map[error]bool{
		RedisNewUserLoginError:true,
		MapKeyError:true,

		EmailMissingErr: true,
		InvalidRecordId: true,
		InvalidEmailErr: true,
		InvalidTitleErr: true,
		InvalidScheduledOnErr: true,
		InvalidDescriptionErr: true,
		InvalidDateTimeFormatErr: true,
		NoUserFoundErr: true,
		NoTaskFoundErr: true,
		NoListFoundErr : true,
		InvalidPayloadError : true,
		NoRecordIdErr : true,
		NoStatusErr : true,
		InvalidStatusChangeErr : true,
		InvalidStatusErr : true,
		AlreadyDeletedRecordErr : true,
		NotInPendingStateErr : true,
		InvalidDateFormatErr : true,

	}
	return errorMessageMap
}

func IsFunctionalError(err error) bool {
	errorMessageMap := GetErrorMessagesMap()

	if (errorMessageMap[err]) {
		return true
	}
	return false
}
