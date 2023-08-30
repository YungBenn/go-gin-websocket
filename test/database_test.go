package test

import (
	"fmt"
	"testing"
	"workshop-web-golang/internal/db"
)

func TestConnectDB(t *testing.T){
	config := &db.ConfigDB{
		Host:     "localhost",
		User:     "postgres",
		Password: "example",
		DBName:   "workshop_len",
		Port:     "5432",
		SSLMode:  "disable",
	}

	_ = db.ConnectDB(config)

	fmt.Println("Test connect to database successfully")
}