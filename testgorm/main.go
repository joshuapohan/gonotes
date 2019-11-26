package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	models "github.com/joshuapohan/testgorm/models"
)

func main() {
	models.OpenDB()
	defer models.CloseDB()
	router := mux.NewRouter()
	router.HandleFunc("/signup", CreateNewUser).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func CreateNewUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user.Email)
	models.DB.Create(&user)
	json.NewEncoder(w).Encode(&user)
}
