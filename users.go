package main

import (
	"fmt"
	"log"
	"miso/users/config"
	"miso/users/controller"
	"net/http"

	"github.com/gorilla/mux"
	_ "gorm.io/driver/postgres"
)

func main() {
	db, err_db := config.Connect()
	if err_db != nil {
		log.Panic(err_db)
		return
	}
	router := mux.NewRouter()
	router.HandleFunc("/users/", controller.CreateUser(db)).Methods("POST")
	router.HandleFunc("/users/auth", controller.GetToken(db)).Methods("POST")
	router.HandleFunc("/users/me", controller.VerifyToken).Methods("GET")
	router.HandleFunc("/users/ping", controller.Ping).Methods("GET")
	http.Handle("/", router)
	fmt.Println("Connected to port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
