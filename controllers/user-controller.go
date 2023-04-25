package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"auth-service/logger"
	. "auth-service/services/user-service"
)

type UserController struct {
	userService IUserService
}

func (u *UserController) Init() {
	u.userService = UserService{}
}

func parseJson[T any](w http.ResponseWriter, body io.ReadCloser, conteiner *T) error {
	err := json.NewDecoder(body).Decode(&conteiner)
	if err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	defer body.Close()
	return nil
}

func writeJson[T any](w http.ResponseWriter, data T) {
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func (userController *UserController) Search(w http.ResponseWriter, r *http.Request) {
	var payload IUserSearchRequest

	err := parseJson(w, r.Body, &payload)
	if err != nil {
		return
	}

	response := userController.userService.Search(payload)
	writeJson(w, response)
}

func (userController *UserController) Create(w http.ResponseWriter, r *http.Request) {
	var payload IUserCreateRequest

	err := parseJson(w, r.Body, &payload)
	if err != nil {
		return
	}
	response := userController.userService.CreateMany(payload)
	writeJson(w, response)
}

func (userController *UserController) Update(w http.ResponseWriter, r *http.Request) {
	var payload IUserUpdateRequest

	err := parseJson(w, r.Body, &payload)
	if err != nil {
		return
	}
	response := userController.userService.UpdateMany(payload)
	writeJson(w, response)
}

func (userController *UserController) Delete(w http.ResponseWriter, r *http.Request) {
	var payload IUserDeleteRequest

	err := parseJson(w, r.Body, &payload)
	if err != nil {
		return
	}
	response := userController.userService.DeleteMany(payload)
	writeJson(w, response)
}
