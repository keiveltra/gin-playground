package main

import (
	"os"
)

func main() {
	args := os.Args

	if len(args) > 1 && (args[1] == "m" || args[1] == "migrate") {
		migrateDatabase()
		return
	}

	if len(args) > 1 && (args[1] == "c" || args[1] == "testConcurrency") {
		testConcurrency()
		return
	}

	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
