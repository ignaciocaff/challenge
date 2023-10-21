package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Env config struct
type EnvApp struct {
	StooqUrl string
	BotQueue string
	RmqHost  string
	RmqPort  string
	RmqUser  string
	RmqPass  string
}

// Get the env configuration
func GetEnv(envFile string) EnvApp {
	err := godotenv.Load(envFile)
	if err != nil {
		fmt.Printf("%v", err)
	}

	return EnvApp{
		StooqUrl: os.Getenv("STOOQ_URL"),
		BotQueue: os.Getenv("BOT_QUEUE"),
		RmqHost:  os.Getenv("RMQ_HOST"),
		RmqPort:  os.Getenv("RMQ_PORT"),
		RmqUser:  os.Getenv("RMQ_USER"),
		RmqPass:  os.Getenv("RMQ_PASS"),
	}
}
