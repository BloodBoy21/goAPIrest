package main

import (
	"encoding/json"
	"log"
	"net/http"
	"server/connect"
	"server/utils"

	"github.com/gorilla/mux"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user_id := vars["id"]
	user := connect.GetUser(user_id)
	message := ""
	status := "OK"
	if user.Id == 0 {
		status = "ERROR"
		message = "User not found"
	}
	response := utils.Response{Status: status, Message: message, Data: user}
	json.NewEncoder(w).Encode(response)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := parseUser(r)

	response := utils.Response{Status: "OK", Message: "User created", Data: connect.CreateUser(user)}
	json.NewEncoder(w).Encode(response)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user_id := vars["id"]
	user := parseUser(r)
	connect.UpdateUser(user_id, user)
	response := utils.Response{Status: "OK", Message: "User updated successfully", Data: connect.UpdateUser(user_id,user)}
	json.NewEncoder(w).Encode(response)
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user_id := vars["id"]
	connect.DeleteUser(user_id)
	response := utils.Response{Status: "OK", Message: "User deleted successfully", Data: utils.User{}}
	json.NewEncoder(w).Encode(response)
}
func parseUser(r *http.Request) utils.User {
	var user utils.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		log.Fatal(err)
	}
	return user
}

func main() {
	connect.InitDB()
	defer connect.CloseConnection()
	r := mux.NewRouter()
	r.HandleFunc("/user/{id}", GetUser).Methods("GET")
	r.HandleFunc("/user", CreateUser).Methods("POST")
	r.HandleFunc("/user/{id}", UpdateUser).Methods("PATCH")
	r.HandleFunc("/user/{id}", DeleteUser).Methods("DELETE")
	log.Println("El servidor se encuentra en el puerto 5000")
	log.Fatal(http.ListenAndServe(":5000", r))
}
