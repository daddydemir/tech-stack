package internal

import (
	"errors"
	"gateway/protogen/golang/users"
)

type DB struct {
	collection []*users.User
}

func NewDB() *DB {
	return &DB{}
}

func (d *DB) AddUser(user *users.User) error {

	for _, u := range d.collection {
		if u.Id == user.Id {
			return errors.New("user already exists")
		}
	}
	d.collection = append(d.collection, user)
	return nil
}

func (d *DB) UpdateUser(user *users.User) error {

	for i, u := range d.collection {
		if u.Id == user.Id {
			d.collection[i] = user
			return nil
		}
	}
	return errors.New("User not found")
}

func (d *DB) GetUser(id int32) (*users.User, error) {

	for _, u := range d.collection {
		if u.Id == id {
			return u, nil
		}
	}
	return nil, errors.New("User not found")
}

func (d *DB) DeleteUser(id int32) error {
	list := make([]*users.User, 0, len(d.collection)-1)
	for _, u := range d.collection {
		if u.Id != id {
			list = append(list, u)
		}
	}
	d.collection = list
	return nil
}
