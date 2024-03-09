package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"modulepath/pkg/models"
	"modulepath/pkg/utils"
	"net/http"
	"strconv"
)

func GetUser(writer http.ResponseWriter, request *http.Request) {
	allUsers := models.GetAllUsers()
	res, err := json.Marshal(allUsers)
	if err != nil {
		http.Error(writer, "Ошибка при обработке данных пользователя", http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	_, err = writer.Write(res)
	if err != nil {
		// В случае ошибки при записи ответа клиенту, можно просто вернуться, поскольку уже был отправлен ответ с ошибкой
		return
	}
}

func AddUser(writer http.ResponseWriter, request *http.Request) {
	addUser := &models.User{}
	utils.ParseBody(request, addUser)
	user := models.AddUser(addUser)
	res, err := json.Marshal(user)
	if err != nil {
		http.Error(writer, "Ошибка при обработке данных пользователя", http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	_, err = writer.Write(res)
	if err != nil {
		// В случае ошибки при записи ответа клиенту, можно просто вернуться, поскольку уже был отправлен ответ с ошибкой
		return
	}
}

func GetUserById(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	userId := vars["userId"]
	id, err := strconv.ParseUint(userId, 0, 0)
	if err != nil {
		log.Println("UserId not integer")
	}
	userDetails, _ := models.GetUserById(id)
	res, err := json.Marshal(userDetails)
	if err != nil {
		http.Error(writer, "Ошибка при обработке данных пользователя", http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	_, err = writer.Write(res)
	if err != nil {
		// В случае ошибки при записи ответа клиенту, можно просто вернуться, поскольку уже был отправлен ответ с ошибкой
		return
	}
}

func UpdateUser(writer http.ResponseWriter, request *http.Request) {
	var updateUser = &models.User{}
	utils.ParseBody(request, updateUser)
	vars := mux.Vars(request)
	userId := vars["userId"]
	id, err := strconv.ParseUint(userId, 0, 0)
	if err != nil {
		log.Println("Userid not integer")
	}
	userDetails, db := models.GetUserById(id)
	if updateUser.Name != "" {
		userDetails.Name = updateUser.Name
	}
	if updateUser.Email != "" {
		userDetails.Email = updateUser.Email
	}
	if updateUser.Address != "" {
		userDetails.Address = updateUser.Address
	}
	db.Save(&userDetails)
	res, err := json.Marshal(userDetails)
	if err != nil {
		http.Error(writer, "Ошибка при обработке данных пользователя", http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	_, err = writer.Write(res)
	if err != nil {
		// В случае ошибки при записи ответа клиенту, можно просто вернуться, поскольку уже был отправлен ответ с ошибкой
		return
	}
}

func DeleteUser(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	userId := vars["userId"]
	id, err := strconv.ParseUint(userId, 0, 0)
	if err != nil {
		log.Println("Userid not integer")
	}
	user := models.DeleteUserById(id)
	res, err := json.Marshal(user)
	if err != nil {
		http.Error(writer, "Ошибка при обработке данных пользователя", http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	_, err = writer.Write(res)
	if err != nil {
		// В случае ошибки при записи ответа клиенту, можно просто вернуться, поскольку уже был отправлен ответ с ошибкой
		return
	}
}
