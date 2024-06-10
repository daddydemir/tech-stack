package internal

import (
	"context"
	"gateway/protogen/golang/users"
	"log"
)

type UserService struct {
	db *DB
	users.UnimplementedUsersServer
}

func NewUserService(db *DB) UserService {
	return UserService{db: db}
}

func (u *UserService) AddUser(_ context.Context, req *users.AddUserReq) (*users.Empty, error) {

	log.Println("AddUser {}", req.User)
	return nil, nil
}

func (u *UserService) UpdateUser(_ context.Context, req *users.UpdateUserReq) (*users.Empty, error) {

	return nil, nil
}

func (u *UserService) GetUser(_ context.Context, req *users.GetUserReq) (*users.GetUserRes, error) {
	return nil, nil
}

func (u *UserService) DeleteUser(_ context.Context, req *users.DeleteUserReq) (*users.Empty, error) {

	return nil, nil
}
