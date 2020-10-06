package db

import (
	"github.com/spf13/viper"
	"gopkg.in/mgo.v2"
	"log"
	"time"
)

func InitMongoDbSession() (mongoSessionMap map[string]*mgo.Session, err error) {

	var dialInfo *mgo.DialInfo
	var mongoSession *mgo.Session

	mongoSessionMap = make(map[string]*mgo.Session)
	mongoSchemas := viper.GetStringMap("mongodb.schema")

	for _, value := range mongoSchemas {
		schemaName := value.(string)
		url := viper.GetStringSlice("database.mongo.schemas." + schemaName + ".url")
		password := viper.GetString("database.mongo.schemas." + schemaName + ".password")
		username := viper.GetString("database.mongo.schemas." + schemaName + ".username")
		poolLimit := viper.GetInt("database.mongo.schemas." + schemaName + ".poolLimit")
		timeOut := viper.GetInt("database.mongo.schemas." + schemaName + ".timeOut")
		authenticationDatabase := viper.GetString("database.mongo.schemas." + schemaName + ".authenticationDatabase")

		if len(username) != 0 {
			dialInfo = &mgo.DialInfo{
				Addrs:     url,
				Username:  username,
				Password:  password,
				PoolLimit: poolLimit,
				Database:  authenticationDatabase,
				Timeout:   time.Duration(timeOut) * time.Second,
			}
		} else {
			dialInfo = &mgo.DialInfo{
				Addrs:     url,
				PoolLimit: poolLimit,
				Timeout:   time.Duration(timeOut) * time.Second,
			}
		}
		mongoSession, err = mgo.DialWithInfo(dialInfo)

		if err != nil {
			log.Println(err)
			return nil, err
		} else {
			mongoSessionMap[schemaName] = mongoSession
		}

	}
	return mongoSessionMap, nil
}
