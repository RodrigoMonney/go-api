package usecases

import "go-api/internal/domain"

type UserUseCase struct {
	repo domain.UserRepository
}

func NewUserUseCase(r domain.UserRepository) *UserUseCase {
	return &UserUseCase{repo: r}
}

func (uc *UserUseCase) GetAllUsers() ([]domain.User, error) {
	return uc.repo.GetAll()
}

func (uc *UserUseCase) GetUserByID(id int) (domain.User, error) {
	return uc.repo.GetByID(id)
}

func (uc *UserUseCase) CreateUser(u domain.User) (domain.User, error) {
	return uc.repo.Create(u)
}
