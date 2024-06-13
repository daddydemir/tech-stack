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

func (u *UserService) AddUser(_ context.Context, req *users.User) (*users.Empty, error) {

	log.Println("AddUser {}", req)
	err := u.db.AddUser(req)
	return nil, err
}

func (u *UserService) UpdateUser(_ context.Context, req *users.User) (*users.Empty, error) {

	err := u.db.UpdateUser(req)
	return nil, err
}

func (u *UserService) GetUser(_ context.Context, req *users.GetUserReq) (*users.User, error) {

	user, err := u.db.GetUser(req.UserId)
	return user, err
}

func (u *UserService) DeleteUser(_ context.Context, req *users.DeleteUserReq) (*users.Empty, error) {

	err := u.db.DeleteUser(req.UserId)

	return nil, err
}
