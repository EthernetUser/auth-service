package services

import (
	"fmt"
	"reflect"

	"auth-service/logger"
)

func (request *IUserSearchRequest) IsEmptyRequest() bool {
	return len(request.Ids) < 1 && (reflect.ValueOf(request.PageSize).IsZero())
}

func (userService UserService) Search(payload IUserSearchRequest) IUserSearchResponse {
	logger.Infof("Got payload %v on userService.Search", payload)

	if payload.IsEmptyRequest() {
		logger.Info("Empty requests")
		return IUserSearchResponse{
			Items: []IUser{},
		}
	}

	result := IUserSearchResponse{
		Items: []IUser{},
	}

	for _, v := range payload.Ids {
		result.Items = append(result.Items, IUser{Id: v, Email: fmt.Sprintf("email%s@gmail.com", v)})
	}

	return result
}

func (userService UserService) CreateMany(payload IUserCreateRequest) IUserCreateResponse {
	return IUserCreateResponse{
		CreatedUsers: []IUser{},
	}
}

func (userService UserService) UpdateMany(payload IUserUpdateRequest) IUserUpdatedResponse {
	return IUserUpdatedResponse{
		UpdatedUsers: []IUser{},
	}
}

func (userService UserService) DeleteMany(ids IUserDeleteRequest) IUserDeleteResponse {
	return IUserDeleteResponse{
		DeletedUserIds: []string{},
	}
}
