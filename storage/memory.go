package storage

import (
	"errors"
	"go-api/models"
)

var users = []models.User{}
var nextID = 1

func GetAllUsers() []models.User {
	return users
}

func GetUserByID(id int) (models.User, error) {
	for _, u := range users {
		if u.ID == id {
			return u, nil
		}
	}
	return models.User{}, errors.New("User not found")
}

func CreateUser(user models.User) models.User {
	user.ID = nextID
	nextID++
	users = append(users, user)
	return user
}