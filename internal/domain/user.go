package domain

type User struct {
	ID    int
	Name  string
	Email string
}

type UserRepository interface {
	GetAll() ([]User, error)
	GetByID(id int) (User, error)
	Create(user User) (User, error)
}
