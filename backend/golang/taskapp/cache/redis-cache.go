package cache

import (
	"encoding/json"
	"fmt"
	"github.com/tb/task-logger/backend/golang/taskapp/models"
	"gopkg.in/redis.v3"
	"time"
)

type redisCache struct {
	host    string
	db      int
	expires time.Duration
}

func NewRedisCache(host string, db int, exp time.Duration) TaskCache {
	return &redisCache{
		host:    "127.0.0.1:6379",
		db:      db,
		expires: exp,
	}
}

func (cache *redisCache) GetClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cache.host,
		Password: "",
		DB:      0,
	})


}

func (cache *redisCache) Set(key string, Task *models.Task) {
	client := cache.GetClient()
	//client.Del("task:*")
	// serialize Task object to JSON
	json, err := json.Marshal(Task)
	if err != nil {
		panic(err)
	}

	_ = client.Set(key, json, cache.expires)


}

func (cache *redisCache) Get(key string) *models.Task {
	client := cache.GetClient()

	val, err := client.Get(key).Result()
	if err != nil {
		return nil
	}

	Task := models.Task{}
	err = json.Unmarshal([]byte(val), &Task)
	if err != nil {
		panic(err)
	}
	return &Task
}


func (cache *redisCache) Del(key string){
	client := cache.GetClient()

	_ = client.Del(key)


}
func (cache *redisCache) Flush(key string) error{
	client := cache.GetClient()
	err := client.FlushAll().Err()

return err
}
func (cache *redisCache) Exists(key string) bool{
	client := cache.GetClient()

	status, _ := client.Exists(key).Result()
	return status


}

func (cache *redisCache) PSubPub (key string)     {
	client := cache.GetClient()

	keyspace := fmt.Sprintf("__keyspace@*__:%s", key)
	ps, _ := client.PSubscribe(keyspace, "expired")

	for {
		fmt.Println("what am i doing here")
		msg, err := ps.ReceiveMessage()
		if msg != nil && err == nil {
			fmt.Printf("Channel[%v] Pattern[%v] Payload[%v]\n", msg.Channel, msg.Pattern, msg.Payload)
		}
		fmt.Println("msg-------->",msg)
	}

}




func (cache *redisCache) List(key string) []models.Task {
	client := cache.GetClient()

	list, err := client.Keys(key).Result()
	if err != nil {

		return nil
	}

	var CacheList []models.Task
	for i := 0; i < len(list); i++ {
		Task := models.Task{}
		listOfCache, err := client.Get(list[i]).Result()
		if err != nil {
			fmt.Println("Error in getting all the keys", err)

		}

		err = json.Unmarshal([]byte(listOfCache), &Task)

		if err != nil {
			fmt.Println("Error in unmarshalling each keys", err)

		}
		CacheList = append(CacheList, Task)

	}

	return CacheList
}
