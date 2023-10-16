package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	// _ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/pressly/goose/v3"
)

func DBConn() (*sql.DB, error) {

	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		log.Fatalf("err loading: %v", err)
	}

	dbDriver := os.Getenv("DB_DRIVER")
	dbUser := os.Getenv("DB_USERNAME")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_DATABASE")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	env := os.Getenv("ENV")

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", dbUser, dbPass, dbName, dbHost, dbPort)
	db, err := sql.Open(dbDriver, connStr)

	_, b, _, _ := runtime.Caller(0)
	// Root folder of this project
	migrateDir := filepath.Join(filepath.Dir(b), "../db/migration")

	if env != "development" {
		if err := goose.Up(db, migrateDir); err != nil {
			return nil, err
		}
	}

	return db, err
}
