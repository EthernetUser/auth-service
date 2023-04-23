package services

type IUser struct {
	Id    string
	Email string
}

type IUserSearchResponse struct {
	Items []IUser
}

type IUserService interface {
	Search(payload IUserSearchRequest) IUserSearchResponse
	CreateMany(payload interface{}) string
	UpdateMany(payload interface{}) string
	DeleteMany(ids []uint) string
}

type UserService struct {
}

type IUserSearchRequest struct {
	Ids      []string
	PageSize int
	Offset   int
}
