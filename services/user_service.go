package services

import (
	"golang_ot/models/mongoModels"
	"golang_ot/repository/mongo"
)

type UserService struct {
	repo *mongo.UserRepository
}

func NewUserService() *UserService {
	return &UserService{repo: mongo.NewUserRepository()}
}

func (s *UserService) CreateUser(user *mongoModels.User) (*mongoModels.User, error) {
	return s.repo.CreateUser(user)
}

func (s *UserService) GetUserByID(id string) (*mongoModels.User, error) {
	return s.repo.GetUserByID(id)
}

// func (s *UserService) UpdateUser(id string, user *mongoModels.User) error {
// 	return s.repo.UpdateUser(id, user)
// }

// func (s *UserService) DeleteUser(id string) error {
// 	return s.repo.DeleteUser(id)
// }
