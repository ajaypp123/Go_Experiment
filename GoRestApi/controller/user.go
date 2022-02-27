package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}

var users map[uint]*User
var count uint

func InitUser() {
	users = make(map[uint]*User)
	count = 0
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, present := mux.Vars(r)["id"]
	if !present {
		json.NewEncoder(w).Encode("Error user not present " + id)
	} else {
		uid, _ := strconv.ParseUint(id, 10, 8)
		var user *User = users[uint(uid)]
		json.NewEncoder(w).Encode(user)
	}
}

func PutUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user *User
	json.NewDecoder(r.Body).Decode(&user)
	user.ID = count
	count++
	users[user.ID] = user
	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	uid, _ := strconv.ParseUint(id, 10, 8)
	var user *User
	json.NewDecoder(r.Body).Decode(&user)
	users[uint(uid)] = user
	json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	uid, _ := strconv.ParseUint(id, 10, 8)
	var user *User = users[uint(uid)]
	delete(users, uint(uid))
	json.NewEncoder(w).Encode(user)
}
