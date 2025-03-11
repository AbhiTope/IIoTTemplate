package repo

type LoginModel struct {
	UserName     string `json:"userName" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterModel struct {
    UserName string `json:"userName" binding:"required"`
    Email string `json:"email" binding:"required"`
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
