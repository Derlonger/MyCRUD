package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/gorilla/mux"
	"modulepath/pkg/models"
	"modulepath/pkg/utils"
)

// GetUser возвращает список всех пользователей
func GetUser(writer http.ResponseWriter, request *http.Request) {
	allUsers := models.GetAllUsers()
	log.Printf("Получены следующие пользователи: %+v\n", allUsers) // Добавлено логирование
	// Проверяем заголовок Accept на предмет запроса JSON
	if request.Header.Get("Accept") == "application/json" {
		// Если запрос на JSON, возвращаем JSON ответ
		log.Println("Запрошен список всех пользователей (JSON).")
		utils.RespondWithJSON(writer, http.StatusOK, allUsers)
		return
	}

	// Если запрос на HTML, возвращаем HTML страницу
	log.Println("Запрошен список всех пользователей (HTML).")
	tmpl, err := template.ParseFiles("templates/users.html")
	if err != nil {
		http.Error(writer, "Ошибка при загрузке шаблона", http.StatusInternalServerError)
		log.Println("Ошибка при загрузке шаблона:", err)
		return
	}

	writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := tmpl.Execute(writer, allUsers); err != nil {
		http.Error(writer, "Ошибка при отображении страницы", http.StatusInternalServerError)
		log.Println("Ошибка при заполнении шаблона:", err)
		return
	}
}

// AddUser добавляет нового пользователя и возвращает его данные в формате JSON
func AddUser(writer http.ResponseWriter, request *http.Request) {
	// Получаем данные из формы
	name := request.FormValue("name")
	email := request.FormValue("email")
	address := request.FormValue("address")

	// Создаем нового пользователя
	newUser := &models.User{
		Name:    name,
		Email:   email,
		Address: address,
	}

	// Добавляем пользователя в базу данных
	user := models.AddUser(newUser)

	// Отправляем ответ клиенту
	res, err := json.Marshal(user)
	if err != nil {
		http.Error(writer, "Ошибка при обработке данных пользователя", http.StatusInternalServerError)
		log.Println("Ошибка при обработке данных пользователя:", err)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	_, err = writer.Write(res)
	if err != nil {
		log.Println("Ошибка при записи ответа клиенту:", err)
		return
	}
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
	log.Printf("Запрошены данные пользователя с ID %d\n", id)
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
	log.Printf("Обновлены данные пользователя с ID %d\n", id)
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
	log.Printf("Удален пользователь с ID %d\n", id)
	utils.RespondWithJSON(writer, http.StatusOK, user)
}
