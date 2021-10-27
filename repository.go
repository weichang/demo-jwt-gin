package main

func insertUser(username string, password string, age uint) (*User, error) {
	user := User{
		Username: username,
		Password: password,
		Age:      age,
	}

	if res := DB.Create(&user); res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}

func findUserByUsername(username string) (*User, error) {
	var user User

	if res := DB.Where("username = ?", username).Find(&user); res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}

func findUserByID(id uint) (*FindUser, error) {
	var user FindUser
	if res := DB.Select("id","username","age").Table("users").Find(&user, id); res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}
