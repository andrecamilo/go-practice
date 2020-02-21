package presenter

import "github.com/andrecamilo/go-practice/clean_tech/ex2/domain/model"

type UserPresenter interface {
	ResponseUsers(u []*model.User) []*model.User
}
