package redis

import (
	"app/config"
	"app/models"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v7"
	"log"
	"reflect"
	"strconv"
	"time"
)

const Delimiter = ":"
const KeyNamePattern = "parkingTaxi" + Delimiter + "%d"

var AllowedIndex = []string{"GlobalID", "ID", "Mode"}

type Redis struct {
	Client *redis.Client
}

func Connect(config config.RedisConfiguration) *Redis {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Hostname + ":" + strconv.Itoa(config.Port),
		Password: config.Password,
		DB:       config.Database,
	})

	_, err := client.Ping().Result()

	if err != nil {
		panic(fmt.Sprintf("i can not connected to Redis. Check config file. %s", err))
	}

	return &Redis{Client: client}
}

func (r Redis) GetPatternName() string {
	return KeyNamePattern
}

func (r Redis) FlushAll() {
	r.Client.FlushAll()
}

func (r Redis) Insert(models []models.ParkingTaxi) {
	pipeline := r.Client.Pipeline()

	IndexData := make(map[string]map[interface{}][]interface{})

	for _, model := range models {
		stringJson, err := json.Marshal(model)
		if err != nil {
			panic(err)
		}

		key := fmt.Sprintf(r.GetPatternName(), model.ID)
		pipeline.Set(key, string(stringJson), 0)

		// https://gist.github.com/drewolson/4771479
		val := reflect.Indirect(reflect.ValueOf(model))
		for i := 0; i < val.NumField(); i++ {
			field := val.Type().Field(i).Name
			value := val.Field(i).Interface()
			if contains(r.GetAllowedIndex(), field) {
				if IndexData[field] == nil {
					IndexData[field] = make(map[interface{}][]interface{})
				}

				IndexData[field][value] = append(IndexData[field][value], model.ID)
			}
		}
	}

	pipeline.Exec()

	r.CreateIndex(IndexData)
}

func (r Redis) FindByQuery(query []string) []models.ParkingTaxi {
	var ids []string
	var keys []string
	var result []models.ParkingTaxi

	start := time.Now()
	ids, _ = r.Client.SInter(query...).Result()

	for _, id := range ids {
		intId, _ := strconv.Atoi(id)
		keys = append(keys, fmt.Sprintf(r.GetPatternName(), intId))
	}
	elapsed := time.Since(start)
	log.Printf("sinter redis func: %s", elapsed)

	startm := time.Now()
	data, _ := r.Client.MGet(keys...).Result()
	elapsedm := time.Since(startm)
	log.Printf("mget redis func: %s", elapsedm)

	startj := time.Now()
	for _, parkingData := range data {
		parkingTaxi := models.ParkingTaxi{}

		b := []byte(parkingData.(string))
		err := json.Unmarshal(b, &parkingTaxi)
		if err != nil {
			panic(err)
		}

		result = append(result, parkingTaxi)
	}
	elapsedj := time.Since(startj)
	log.Printf("json redis func: %s", elapsedj)

	return result
}

func (r Redis) CreateIndex(IndexData map[string]map[interface{}][]interface{}) {
	pipeline := r.Client.Pipeline()
	for field, data := range IndexData {
		for valueIndex, id := range data {
			pipeline.SAdd(fmt.Sprintf("%s:%v", field, valueIndex), id...)
		}
	}

	_, err := pipeline.Exec()
	if err != nil {
		panic(err)
	}
}

func (r Redis) GetAllowedIndex() []string {
	return AllowedIndex
}

func (r Redis) GetDB() *redis.Client {
	return r.Client
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
