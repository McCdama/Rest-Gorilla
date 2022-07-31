package main

import (
	"fmt"
	"os"

	"github.com/McCdama/Rest-Gorilla/app"
)

func main() {
	fmt.Print("Starting")
	a := app.App{}
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	a.Run(":8083")

}
