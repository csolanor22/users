package main

import (
	"bytes"
	"miso/users/controller"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Test_CreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDB, _ := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})

	rows := sqlmock.NewRows([]string{"id", "username", "password", "email", "expire_at", "created_at", "token", "salt"}).
		AddRow(1, "John Doe", "e0d803e4f6bd9621e2bf87b359c700dcd839b0534d2421f179d62f53c628b216573d6ddaaf8465ded5a2d78bb89a4a3c0c27b71163de0707ece3336516fb1ced", "mail1@mail.com", nil, nil, nil, "lgTeMaPEZQ")

	mock.ExpectQuery("SELECT(.*)").WithArgs("John Doe", "mock@mail.com").WillReturnRows(rows)

	r := mux.NewRouter()
	r.HandleFunc("/users", controller.CreateUser(gormDB))

	var jsonStr = []byte(`{"username": "John Doe", "Password": "contrasena", "email": "mock@mail.com" }`)
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	checkResponseCode(t, http.StatusPreconditionFailed, w.Code)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
