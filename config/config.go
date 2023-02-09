package config

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	localhost := os.Getenv("HOST_DB")
	if len(localhost) == 0 {
		localhost = "localhost"
	}
	user := os.Getenv("USER_DB")
	if len(user) == 0 {
		user = "postgres"
	}
	password := os.Getenv("PASS_DB")
	if len(password) == 0 {
		password = "password"
	}
	port := os.Getenv("PORT_DB")
	if len(port) == 0 {
		port = "5432"
	}
	dbname := os.Getenv("NAME_DB")
	if len(dbname) == 0 {
		dbname = "test"
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Bogota", localhost, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	return db, nil
}
