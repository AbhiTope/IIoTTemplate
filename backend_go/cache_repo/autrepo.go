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
    Role string `json:"role"`
    Password string `json:"password" binding:"required"`
    IsLocked bool
}

var users []RegisterModel 

func (r RegisterModel) Add() error {
	users = append(users, r)
	return nil 
}

func GetActiveUsers() []RegisterModel {
	var result []RegisterModel
	for _, user := range users {
		if !user.IsLocked {
			user.Password = ""
			result = append(result, user)
		}
	}
    return result
}

func GetAllUsers() []RegisterModel {
	var result []RegisterModel
	for _, user := range users {
			user.Password = ""
			result = append(result, user)
	}
    return result
}

func GetUser(userName string) (RegisterModel, error){

	for _, user := range users {
		if user.UserName == userName {
			user.Password = ""
			return user, nil
		}
	}
	var result RegisterModel

	return result, errors.New("user not found")

}

func GetPassword(userName string) (string, error){

	for _, user := range users {
		if user.UserName == userName {
			if user.IsLocked{
				return "", errors.New("user is locked")
			}
			return user.Password, nil
		}
	}

	return "", errors.New("user not found")

}

