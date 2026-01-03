package cache_redis_config

import (
	"fmt"
	"os"
	"strconv"
)

func GetRedisPortOrDefault() int {
	redisPort := os.Getenv("REDIS_PORT")
	if redisPort == "" {
		return 6379
	}
	port, err := strconv.Atoi(redisPort)
	if err != nil {
		fmt.Printf("Error parsing REDIS_PORT: %s. Defaulting to 6379.\n", err)
		return 6379
	}
	return port
}