package services

import (
	"auth-service/logger"
	"fmt"
	"reflect"
)

type UserService struct {
}

type ISearchRequest struct {
	Ids      []string
	PageSize int
	Offset   int
}

func (request *ISearchRequest) IsEmptyRequest() bool {
	return len(request.Ids) < 1 && (reflect.ValueOf(request.PageSize).IsZero())
}

type IUser struct {
	Id    string
	Email string
}

type ISearchResponse struct {
	Items []IUser
}

type IUserService interface {
	Search(payload ISearchRequest) ISearchResponse
	CreateMany(payload interface{}) string
	UpdateMany(payload interface{}) string
	DeleteMany(ids []uint) string
}

func (userService UserService) Search(payload ISearchRequest) ISearchResponse {
	logger.Infof("Got payload %v on userService.Search", payload)

	if payload.IsEmptyRequest() {
		logger.Info("Empty requests")
		return ISearchResponse{
			Items: []IUser{},
		}
	}

	result := ISearchResponse{
		Items: []IUser{},
	}

	for _, v := range payload.Ids {
		result.Items = append(result.Items, IUser{Id: v, Email: fmt.Sprintf("email%s@gmail.com", v)})
	}

	return result
}

func (userService UserService) CreateMany(payload interface{}) string {

	return ""
}

func (userService UserService) UpdateMany(payload interface{}) string {

	return ""
}

func (userService UserService) DeleteMany(ids []uint) string {

	return ""
}
