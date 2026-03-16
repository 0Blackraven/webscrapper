package utils

import (
	"fmt"
	"strings"

	"github.com/0Blackraven/webscrapper/utils/redis"
)

func AddJob(url string, depth int) {
    normalizedUrl := cleanURL(url)
    if normalizedUrl == "" {
        return
    }

    isNew:= redis.Add(normalizedUrl) 
    if isNew {
        redis.AddToQueue(redis.Job{Url: normalizedUrl, Depth: depth})
    } else {

        fmt.Printf("[Duplicate] Skipping %s\n", normalizedUrl)
    }
}

func cleanURL(url string) string {
    u := strings.TrimSpace(url)
    u = strings.ToLower(u) 
    u = strings.TrimSuffix(u, "/")
    u = strings.Replace(u, "https://", "http://", 1) 
    return u
}

func GetJob() string {
	job := redis.RemoveFromQueue()
	return job.Url
}