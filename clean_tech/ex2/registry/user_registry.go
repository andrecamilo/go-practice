package registry

import (
	"github.com/andrecamilo/go-practice/clean_tech/ex2/interface/controller"
	ip "github.com/andrecamilo/go-practice/clean_tech/ex2/interface/presenter"
	ir "github.com/andrecamilo/go-practice/clean_tech/ex2/interface/repository"
	"github.com/andrecamilo/go-practice/clean_tech/ex2/usecase/interactor"
	up "github.com/andrecamilo/go-practice/clean_tech/ex2/usecase/presenter"
	ur "github.com/andrecamilo/go-practice/clean_tech/ex2/usecase/repository"
)

func (r *registry) NewUserController() controller.UserController {
	return controller.NewUserController(r.NewUserInteractor())
}

func (r *registry) NewUserInteractor() interactor.UserInteractor {
	return interactor.NewUserInteractor(r.NewUserRepository(), r.NewUserPresenter())
}

func (r *registry) NewUserRepository() ur.UserRepository {
	return ir.NewUserRepository(r.db)
}

func (r *registry) NewUserPresenter() up.UserPresenter {
	return ip.NewUserPresenter()
}
