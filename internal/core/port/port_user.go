package port

import (
	"samsamoohooh-mini-api/internal/core/domain"
	"samsamoohooh-mini-api/internal/core/dto"
)

type UserService interface {
	CreateUser(userDto *dto.CreateUserRequest) (*dto.CreateUserResponse, error)
	GetUser(id int) (*dto.GetUserResponse, error)
	UpdateUser(id int, dto *dto.UpdateUserRequest) (*dto.UpdateUserResponse, error)
	DeleteUser(id int) (*dto.DeleteUserResponse, error)
}

type UserRepository interface {
	CreateUser(user *domain.User) (*domain.User, error)
	GetUserByID(id int) (*domain.User, error)
	UpdateUser(id int, user *domain.User) (*domain.User, error)
	DeleteUser(id int) (*domain.User, error)
}
