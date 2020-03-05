package repository

import "github.com/andrecamilo/go-practice/clean_tech/ex2/domain/model"

type UserRepository interface {
	FindAll(u []*model.User) ([]*model.User, error)
}
