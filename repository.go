package main

import "github.com/weichang/demo-jwt-gin/model"

func insertUser(username string, password string, age uint) (*model.User, error) {
	user := model.User{
		Username: username,
		Password: password,
		Age:      age,
	}

	if res := DB.Create(&user); res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}

func findUserByUsername(username string) (*model.User, error) {
	var user model.User

	if res := DB.Where("username = ?", username).Find(&user); res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}

func findUserByID(id uint) (*model.FindUser, error) {
	var user model.FindUser
	if res := DB.Select("id","username","age").Table("users").Find(&user, id); res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}
