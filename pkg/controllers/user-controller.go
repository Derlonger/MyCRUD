package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"modulepath/pkg/models"
	"modulepath/pkg/utils"
)

// GetUser возвращает всех пользователей в формате JSON
func GetUser(writer http.ResponseWriter, _ *http.Request) {
	allUsers := models.GetAllUsers()
	utils.RespondWithJSON(writer, http.StatusOK, allUsers)
}

// AddUser добавляет нового пользователя и возвращает его данные в формате JSON
func AddUser(writer http.ResponseWriter, request *http.Request) {
	addUser := &models.User{}
	utils.ParseBody(request, addUser)
	user := models.AddUser(addUser)
	utils.RespondWithJSON(writer, http.StatusOK, user)
}

// GetUserById возвращает данные пользователя по указанному ID в формате JSON
func GetUserById(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	userId := vars["userId"]
	id, err := strconv.ParseUint(userId, 0, 0)
	if err != nil {
		log.Println("UserId not integer")
	}
	userDetails, _ := models.GetUserById(id)
	utils.RespondWithJSON(writer, http.StatusOK, userDetails)
}

// UpdateUser обновляет данные пользователя по указанному ID и возвращает обновленные данные в формате JSON
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
	utils.RespondWithJSON(writer, http.StatusOK, userDetails)
}

// DeleteUser удаляет пользователя по указанному ID и возвращает результат операции в формате JSON
func DeleteUser(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	userId := vars["userId"]
	id, err := strconv.ParseUint(userId, 0, 0)
	if err != nil {
		log.Println("Userid not integer")
	}
	user := models.DeleteUserById(id)
	utils.RespondWithJSON(writer, http.StatusOK, user)
}
