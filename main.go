package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func HelloWorld() string {
	return os.Getenv("MESSAGE")
}

func main() {
	godotenv.Load()

	fmt.Println(HelloWorld())
}
