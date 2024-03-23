package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvMongoURI() string {
	err := godotenv.Load()

	if err != nil {
		log.Fatalln("error .env")
	}

	mongoURI := os.Getenv("MONGOURI")
	return mongoURI
}
