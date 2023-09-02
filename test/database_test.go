package test

import (
	"log"
	"testing"
	"workshop-web-golang/config"
	"workshop-web-golang/internal/db"
)

func TestConnectDB(t *testing.T) {
	env, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %s", err)
	}

	config := &db.ConfigDB{
		Host:     env.DB_HOST,
		User:     env.DB_USER,
		Password: env.DB_PASS,
		DBName:   env.DB_NAME,
		Port:     env.DB_PORT,
		SSLMode:  env.DB_SSLMODE,
	}

	if config.Host == "localhost" ||
		config.User == "postgres" ||
		config.Password == "example" ||
		config.DBName == "workshop_len" ||
		config.Port == "5432" ||
		config.SSLMode == "disable" {
		t.Error("Database config didn't work")
	}
}
