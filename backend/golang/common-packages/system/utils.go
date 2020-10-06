package system

import (
	"github.com/spf13/viper"
	"time"
	"strings"
	"reflect"
)

func GetDatabaseName(collectionName string) (databaseName string) {

	database := viper.GetString("mongodb.collection." + collectionName + ".db")
	if len(database) > 0 {
		return viper.GetString("mongodb.schema." + database)
	} else {
		return viper.GetString("mongodb.defaultSchema")
	}

}

func Contains(val interface{}, arr interface{}) (exists bool, index int) {
	exists = false
	index = -1
	switch reflect.TypeOf(arr).Kind() {
	case reflect.Array:
		s:= reflect.ValueOf(arr)
		for i:=0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) {
				index = i
				exists = true
				return
			}
		}

	}
	return
}

func ChangeStringToDateTime(dateTimeStr string) (time.Time, error) {
	date_time_arr := strings.Split(dateTimeStr, " ")
	dateParseFormat := date_time_arr[0]+"T"+date_time_arr[1]+":00Z"
	parsedDateTime, err := time.Parse(time.RFC3339,dateParseFormat);
	if err != nil {
		return time.Now(), InvalidDateTimeFormatErr
	}else {
		return parsedDateTime, nil
	}
}