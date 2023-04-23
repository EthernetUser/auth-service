package controllers

import (
	"auth-service/logger"
	. "auth-service/services/user-service"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

func (userController *UserController) Search(w http.ResponseWriter, r *http.Request) {
	var payload IUserSearchRequest

	err := parseJson(w, r.Body, &payload)
	if err != nil {
		return
	}

	response := userController.userService.Search(payload)
	json.NewEncoder(w).Encode(response)
}

func (u *UserController) Create(w http.ResponseWriter, r *http.Request) {
	var payload IUserSearchRequest

	err := parseJson(w, r.Body, &payload)
	if err != nil {
		return
	}
	w.Write(fmt.Append([]byte{}, "Create User"))
}

func (u *UserController) Update(w http.ResponseWriter, r *http.Request) {
	var payload IUserSearchRequest

	err := parseJson(w, r.Body, &payload)
	if err != nil {
		return
	}
	w.Write(fmt.Append([]byte{}, "Update User"))
}

func (u UserController) Delete(w http.ResponseWriter, r *http.Request) {
	var payload IUserSearchRequest

	err := parseJson(w, r.Body, &payload)
	if err != nil {
		return
	}
	w.Write(fmt.Append([]byte{}, "Delete User"))
}
