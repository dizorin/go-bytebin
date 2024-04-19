package utils

import (
	"github.com/gofiber/fiber/v2/log"
	"os"
	"strconv"
	"time"
)

func GetenvBool(key string) bool {
	parseEnv, err := strconv.ParseBool(os.Getenv(key))
	if err != nil {
		log.Warn("Parse bool: key=", key, " err=", err)
		parseEnv = false
	}
	return parseEnv
}

func GetenvDuration(key string) time.Duration {
	parseEnv, err := time.ParseDuration(os.Getenv(key))
	if err != nil {
		log.Warn("Parse duration: key=", key, " err=", err)
		parseEnv = time.Minute
	}
	return parseEnv
}
