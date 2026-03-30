package redis

import (
	"fmt"
)

func init() {
	redisdb.Do(ctx, "DEL", "mybloom")
	redisdb.Do(ctx, "BF.RESERVE", "mybloom", 0.01, 1000000)
}

func Add(url string) bool {
	result, err := redisdb.Do(ctx, "BF.ADD", "mybloom", url).Bool()
	if err != nil {
		fmt.Printf("Accounted error while trying to add url to bloomfilter\nError: [%v]\n", err)
		return false
	}
	return result
}

func Check(url string) bool {
	result, err := redisdb.Do(ctx, "BF.EXISTS", "mybloom", url).Bool()
	if err != nil {
		fmt.Printf("Accounted error while trying to find url in bloomfilter\nError: [%v]\n", err)
		return false
	}
	return result
}