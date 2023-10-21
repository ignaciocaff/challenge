package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

const DEFAULT_PORT_IF_EMPTY = "3000"
const DEFAULT_WSPORT_IF_EMPTY = "3001"

// Env config struct
type EnvApp struct {
	Port        string
	JwtSecret   string
	GinMode     string
	BotQueue    string
	MongoHost   string
	MongoPort   string
	MongoUser   string
	MongoPass   string
	MongoDbName string
	RmqHost     string
	RmqPort     string
	RmqUser     string
	RmqPass     string
}

// Get the env configuration
func GetEnv(envFile string) EnvApp {
	err := godotenv.Load(envFile)
	if err != nil {
		fmt.Printf("%v", err)
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = DEFAULT_PORT_IF_EMPTY
	}

	return EnvApp{
		Port:        port,
		JwtSecret:   os.Getenv("JWT_SECRET"),
		GinMode:     os.Getenv("GIN_MODE"),
		BotQueue:    os.Getenv("BOT_QUEUE"),
		MongoHost:   os.Getenv("MONGO_HOST"),
		MongoPort:   os.Getenv("MONGO_PORT"),
		MongoUser:   os.Getenv("MONGO_USER"),
		MongoPass:   os.Getenv("MONGO_PASS"),
		MongoDbName: os.Getenv("MONGO_DB_NAME"),
		RmqHost:     os.Getenv("RMQ_HOST"),
		RmqPort:     os.Getenv("RMQ_PORT"),
		RmqUser:     os.Getenv("RMQ_USER"),
		RmqPass:     os.Getenv("RMQ_PASS"),
	}
}
