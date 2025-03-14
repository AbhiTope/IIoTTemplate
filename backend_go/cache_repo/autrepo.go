package repo

import (
	"errors"
)

type LoginModel struct {
	UserName     string `json:"userName" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterModel struct {
    UserName string `json:"userName" binding:"required"`
    FirstName string `json:"firstName" binding:"required"`
    LastName string `json:"lastName" binding:"required"`
    Email string `json:"email" binding:"required"`
    Role string `json:"role" binding:"required"`
    Password string `json:"password" binding:"required"`
}

var users []RegisterModel 

func (r LoginModel) Validate() bool {

	for _, user := range users {
		if user.UserName == r.UserName && user.Password == r.Password {
			return true
		}
	}

    return false

}

func (r RegisterModel) Add() error {
    users = append(users, r)
    return nil 
}

func GetUsers() []RegisterModel {
    return users
}

func GetUser(userName string) (RegisterModel, error){

	for _, user := range users {
		if user.UserName == userName {
			return user, nil
		}
	}
	var result RegisterModel

	return result, errors.New("user now found")

}

