package main

import (
	"fmt"
	"go-learn/router"
	"log"
	"net/http"
	"os"
)

func main() {
	handler := router.New()
	server := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%s", os.Getenv("PORT")),
		Handler: handler,
	}
	log.Println("server running at", server.Addr)

	// // Run Goose migrations
	// if err := runGooseUp(); err != nil {
	// 	log.Fatalf("Failed to run Goose migrations: %v", err)
	// }

	server.ListenAndServe()

}

// func runGooseUp() error {
// 	log.Println("Goose migrations....")
// 	// Get the current working directory
// 	cwd, err := os.Getwd()
// 	if err != nil {
// 		return err
// 	}

// 	// Set up the goose command to run migrations
// 	cmd := exec.Command("goose", "-dir", fmt.Sprintf("%s/db/migration", cwd), "postgres", os.Getenv("DB_URL"), "up")

// 	// Set the appropriate environment variables for the goose command
// 	cmd.Env = os.Environ()

// 	// Run the goose command
// 	output, err := cmd.CombinedOutput()
// 	if err != nil {
// 		return fmt.Errorf("failed to run goose up: %v, output: %s", err, string(output))
// 	}

// 	log.Println("Goose migrations ran successfully.")
// 	return nil
// }
