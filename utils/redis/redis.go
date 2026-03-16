package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

var redisdb *redis.Client
var ctx = context.Background()

func init() {
	godotenv.Load()
	var addr = os.Getenv("REDIS_ADDR")
	var pass = os.Getenv("REDIS_PASS")
	var options = redis.Options{
		Addr:     addr,
		Password: pass,
		DB:       0,
	}
	redisdb = redis.NewClient(&options)
	// defer redisdb.Close() me is just dumb
}

type Job struct {
	Url   string
	Depth int  
}

func AddToQueue(job Job) {
	data, err := json.Marshal(job)
	if err != nil {
		fmt.Printf("Error marshaling job: [%v]\n", err)
		return
	}
	redisdb.LPush(ctx, "urlQueue", data)
}

func RemoveFromQueue() Job {
	val, err := redisdb.RPop(ctx, "urlQueue").Result()
	if err != nil {
		if err != redis.Nil {
			fmt.Printf("Error while trying to remove url from queue\nError: [%v]\n", err)
		}
		return Job{}
	}

	var job Job
	err = json.Unmarshal([]byte(val), &job)
	if err != nil {
		fmt.Printf("Error unmarshaling job: [%v]\n", err)
		return Job{}
	}
	return job
}
