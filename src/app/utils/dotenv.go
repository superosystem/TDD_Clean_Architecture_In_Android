package utils

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func ReadEnvFile() {
	err := godotenv.Load(".env")

	if err != nil {
		logrus.Error("Error loading env file!")
		panic(err)
	}
}
