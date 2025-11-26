// Package helpers provides a set of utility functions.
package helpers

import (
	"cache-redis-config/config"
	"fmt"
	"os"
)

// GetConfig returns the config instance.
func GetConfig() *config.Config {
	return config.GetInstance()
}

// GetRedisConfig returns the Redis config instance.
func GetRedisConfig() *config.RedisConfig {
	return GetConfig().Redis
}

// GetLogger returns the logger instance.
func GetLogger() *Logger {
	return config.GetInstance().Logger
}

// InitLogger initializes the logger based on the log level from the config.
func InitLogger() {
	logLevel := GetConfig().Log.Level
	switch logLevel {
	case config.LogLevelDebug:
		log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	case config.LogLevelInfo:
		log.SetFlags(log.Ltime)
	case config.LogLevelError:
		log.SetFlags(0)
	}
}

// GetRedisPrefix returns the Redis prefix.
func GetRedisPrefix() string {
	return GetConfig().Redis.Prefix
}

// GetDatabaseName returns the database name.
func GetDatabaseName() string {
	return GetConfig().Database.Name
}

// GetDatabaseUser returns the database user.
func GetDatabaseUser() string {
	return GetConfig().Database.User
}

// GetDatabasePassword returns the database password.
func GetDatabasePassword() string {
	return GetConfig().Database.Password
}

// GetDatabasePort returns the database port.
func GetDatabasePort() int {
	return GetConfig().Database.Port
}

// GetDatabaseHost returns the database host.
func GetDatabaseHost() string {
	return GetConfig().Database.Host
}

// GetDatabaseSSL returns whether SSL is enabled or not.
func GetDatabaseSSL() bool {
	return GetConfig().Database.Ssl
}

// GetDatabaseTimeout returns the database timeout.
func GetDatabaseTimeout() time.Duration {
	return GetConfig().Database.Timeout
}

// GetRedisTimeout returns the Redis timeout.
func GetRedisTimeout() time.Duration {
	return GetConfig().Redis.Timeout
}

// GetCacheKey returns the cache key.
func GetCacheKey() string {
	return GetConfig().Cache.Key
}

// GetCacheTTL returns the cache TTL.
func GetCacheTTL() time.Duration {
	return GetConfig().Cache.TTL
}

// GetRedisPassword returns the Redis password.
func GetRedisPassword() string {
	return GetConfig().Redis.Password
}

// GetRedisHost returns the Redis host.
func GetRedisHost() string {
	return GetConfig().Redis.Host
}

// GetRedisPort returns the Redis port.
func GetRedisPort() int {
	return GetConfig().Redis.Port
}

// GetRedisDatabase returns the Redis database index.
func GetRedisDatabase() int {
	return GetConfig().Redis.Database
}

// GetRedisPassword returns the Redis password.
func GetRedisPassword() string {
	return GetConfig().Redis.Password
}

// GetRedisTimeout returns the Redis timeout.
func GetRedisTimeout() time.Duration {
	return GetConfig().Redis.Timeout
}

// GetRedisEnableKeepAlive returns whether Redis keep alive is enabled or not.
func GetRedisEnableKeepAlive() bool {
	return GetConfig().Redis.EnableKeepAlive
}

// GetRedisKeepAliveInterval returns the Redis keep alive interval.
func GetRedisKeepAliveInterval() time.Duration {
	return GetConfig().Redis.KeepAliveInterval
}

// GetRedisKeepAliveTimeout returns the Redis keep alive timeout.
func GetRedisKeepAliveTimeout() time.Duration {
	return GetConfig().Redis.KeepAliveTimeout
}

// GetRedisSaveNetworkTraffic returns whether Redis save network traffic is enabled or not.
func GetRedisSaveNetworkTraffic() bool {
	return GetConfig().Redis.SaveNetworkTraffic
}