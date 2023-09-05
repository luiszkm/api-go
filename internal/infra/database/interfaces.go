package database

import "github.com/luiszkm/api/internal/Domain/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}