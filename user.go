package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []User
	DB.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idClean, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if idClean == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var user User
	if err := DB.First(&user, params["id"]).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	if user.Name == "" || user.Address == "" || user.Age == 0 || user.Email == "" || user.Phone == "" || user.Username == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	DB.Create(&user)
	json.NewEncoder(w).Encode(user)
}

func UpdateUserByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idClean, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if idClean == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var user User
	DB.First(&user, params["id"])
	json.NewDecoder(r.Body).Decode(&user)
	if user.Name == "" || user.Address == "" || user.Age == 0 || user.Email == "" || user.Phone == "" || user.Username == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	DB.Save(&user)
	json.NewEncoder(w).Encode(user)
}

func DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idClean, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if idClean == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	var user User
	DB.Delete(&user, params["id"])
}
