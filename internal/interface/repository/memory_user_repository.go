package repository

import (
	"errors"
	"go-api/internal/domain"
)

type memoryUserRepo struct {
	users  []domain.User
	nextID int
}

func NewMemoryUserRepo() domain.UserRepository {
	return &memoryUserRepo{nextID: 1}
}

func (r *memoryUserRepo) GetAll() ([]domain.User, error) {
	if len(r.users) == 0 {
		return []domain.User{}, nil
	}

	return r.users, nil
}

func (r *memoryUserRepo) GetByID(id int) (domain.User, error) {
	for _, u := range r.users {
		if u.ID == id {
			return u, nil
		}
	}
	return domain.User{}, errors.New("user not found")
}

func (r *memoryUserRepo) Create(user domain.User) (domain.User, error) {
	user.ID = r.nextID
	r.nextID++
	r.users = append(r.users, user)
	return user, nil
}
