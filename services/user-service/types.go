package services

type IUser struct {
	Id    string
	Email string
}

type IUserService interface {
	Search(payload IUserSearchRequest) IUserSearchResponse
	CreateMany(payload IUserCreateRequest) IUserCreateResponse
	UpdateMany(payload IUserUpdateRequest) IUserUpdatedResponse
	DeleteMany(payload IUserDeleteRequest) IUserDeleteResponse
}

type UserService struct{}

type IUserSearchRequest struct {
	Ids      []string
	PageSize int
	Offset   int
}

type IUserSearchResponse struct {
	Items []IUser
}

type IUserCreateRequest struct {
	Users []IUser
}

type IUserCreateResponse struct {
	CreatedUsers []IUser
}

type IUserUpdateRequest struct {
	Users []IUser
}

type IUserUpdatedResponse struct {
	UpdatedUsers []IUser
}

type IUserDeleteRequest struct {
	Ids []string
}

type IUserDeleteResponse struct {
	DeletedUserIds []string
}
