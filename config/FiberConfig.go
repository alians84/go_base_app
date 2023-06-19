package config

import (
	"github.com/gofiber/fiber/v2"
	"os"
	"strconv"
	"time"
)

func FiberConfig() fiber.Config {
	readTimeout, _ := strconv.Atoi(os.Getenv("SERVER_READ_TIMEOUT"))

	return fiber.Config{
		ReadTimeout: time.Duration(readTimeout) * time.Second,
	}
}
