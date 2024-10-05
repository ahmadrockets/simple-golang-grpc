package repository

import (
	"context"
	"go-server/model"
)

type UserRepositoryInterface interface {
	GetAllUser(ctx context.Context) (res []model.User, err error)
	Register(ctx context.Context, name, address, phone, email string) (res model.User, err error)
}

type UserRepository struct{}

func NewUserRepository() UserRepositoryInterface {
	return &UserRepository{}
}

func (r *UserRepository) GetAllUser(ctx context.Context) (res []model.User, err error) {
	res = append(res, model.User{
		Name:    "Ahmad Fahrudin",
		Address: "Gunungkidul",
		Phone:   "6282226566222",
		Email:   "ahmad@email.com",
	})
	err = nil

	return
}

func (s *UserRepository) Register(ctx context.Context, name, address, phone, email string) (res model.User, err error) {
	res = model.User{
		Name:    name,
		Address: address,
		Phone:   phone,
		Email:   email,
	}
	return
}
