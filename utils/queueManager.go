package utils

import (
	"strings"

	"github.com/0Blackraven/webscrapper/utils/redis"
)

func AddJob(url string, depth int) {
	normalizedUrl := url
	if strings.HasPrefix(normalizedUrl, "https://") {
		normalizedUrl = "http://" + normalizedUrl[8:]
	}
	normalizedUrl = strings.TrimSuffix(normalizedUrl, "/")
	job := redis.Job{
		Url:   normalizedUrl,
		Depth: depth,
	}
	if !redis.Check(normalizedUrl) {
		redis.AddToQueue(job)
	}
}

func GetJob() string {
	job := redis.RemoveFromQueue()
	redis.Add(job.Url)
	return job.Url
}