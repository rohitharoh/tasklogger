package Validationspackage

import (
	"regexp"
	"github.com/sirupsen/logrus"
	"github.com/tb/task-logger/backend/golang/taskapp/models"


	"github.com/tb/task-logger/backend/golang/common-packages/system"

	"strings"
	"strconv"
	_"flag"
	"time"
)

func ValidateEmail(email string) bool {
	//Re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

	//re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	//Re := regexp.MustCompile(`^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,3}))$`)
	Re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-z]{2,4}$`)
	return Re.MatchString(email)
}


func ValidateName(name string) bool {
	Re := regexp.MustCompile(`^[a-zA-Z\s?\.]+$`)
	return Re.MatchString(name)
}

func ValidateDateFormat(dateToValidate string) bool {

	if len(dateToValidate) == 16 {
		date_time_arr := strings.Split(dateToValidate, " ")
		if len(date_time_arr) == 2 {
			datePart := date_time_arr[0]
			timePart := date_time_arr[1]
			if len(datePart) != 10 || len(timePart) != 5 {
				return false
			}else {
				dateParts := strings.Split(datePart, "-")
				if len(dateParts) != 3 {
					return false
				}
				timeParts := strings.Split(timePart, ":")
				if len(timeParts) != 2 {
					return false
				}
				if len(dateParts[0]) != 4 || len(dateParts[1]) != 2 || len(dateParts[2]) != 2 || len(timeParts[0]) != 2 || len(timeParts[1]) != 2 {
					return false
				}

				providedYear, err := strconv.Atoi(dateParts[0])
				if err != nil {
					return false
				}
				if providedYear < 1970 {
					return false
				}

				providedMonth, err := strconv.Atoi(dateParts[1])
				if err != nil {
					return false
				}
				if (providedMonth < 00 || providedMonth > 12) {
					return false
				}

				providedDay, err := strconv.Atoi(dateParts[2])
				if err != nil {
					return false
				}

				if (providedDay < 00 || providedDay > 31 ) {
					return false
				}
				providedHour, err := strconv.Atoi(timeParts[0])
				if err != nil {
					return false
				}
				if (providedHour < 00 || providedHour > 23 ) {
					return false
				}

				providedMin, err := strconv.Atoi(timeParts[1])
				if err != nil {
					return false
				}
				if (providedMin < 00 || providedMin > 59 ) {
					return false
				}
				dateParseFormat := datePart+"T"+timePart+":00Z"
				_, err = time.Parse(time.RFC3339,dateParseFormat);
				if err != nil {
					return false
				}
				return true
			}


		} else {
			return false
		}
	}else {
		return false
	}


	/*l := len(models.AddTaskInput("ScheduledOn"))

	for l > 0 {
		l--
		if !unicode.IsSpace(l) {


			return l[0]
		}
	}


	//test := strings.Fields(scheduledOn)




	s := strings.SplitN(scheduledOn, "", 10)

        date := strings.Split(s, "-")
	for _, v := range date {
		




		//parsing date and time
	l, err := time.Parse("2006-01-02 15:04", "2011-01-19 22:15")
	if err != nil {
		fmt.Println(err)
		return l, nil
	}*/



}


/*func ValidateDate(name string) bool {
	Re := regexp.MustCompile(`^[0-9]+\-[0-9]+\-[0-9]`)
	return Re.MatchString(name)
}

*/
func ValidateAddTaskInput(logger *logrus.Entry, taskInputInfo models.AddTaskInput)(error) {

	if taskInputInfo.Title == "" {
		return system.InvalidTitleErr

	} else if taskInputInfo.ScheduledOn == "" {

		return system.InvalidScheduledOnErr

	} else if taskInputInfo.Description == "" {

		return system.InvalidDescriptionErr

	} /*else if !ValidateDateFormat(taskInputInfo.ScheduledOn) {
		return system.InvalidDateTimeFormatErr
	}*/
	_, err := time.Parse(system.CUSTOM_DATE_FORMAT, taskInputInfo.ScheduledOn)
	if err != nil {
		return system.InvalidDateFormatErr
	}
	return nil
}

