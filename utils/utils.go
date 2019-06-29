package utils

import (
	"fmt"
	"os"

	env "github.com/joho/godotenv"
)

// LoadEnv should load .env file
func LoadEnv() {
	env.Load()

	fmt.Println(os.Getenv("PORT"))
}
