package redis

import (

	"fmt"
	"os"
    "github.com/joho/godotenv"
	"github.com/RedisBloom/redisbloom-go"
)

var bloomFilter *redis_bloom_go.Client

func init() {
    godotenv.Load()
    var pass = os.Getenv("REDIS_PASS")
    var addr = os.Getenv("REDIS_ADDR")
    bloomFilter = redis_bloom_go.NewClient(
        addr, "bloom", &pass,
    )
}

func Add(url string) {
    _, err := bloomFilter.Add("mybloom", url)
    if err != nil {
        fmt.Printf("Accounted error while trying to add url to bloomfilter\nError: [%v]\n", err)
    }
}

func Check(url string) bool{
    var ok = false
    ok,err := bloomFilter.Exists("mybloom", url)
    if err != nil {
        fmt.Printf("Accounted error while trying to find url in bloomfilter\nError: [%v]\n",err)
    }
    return ok
}