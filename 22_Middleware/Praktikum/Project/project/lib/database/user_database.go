package database

import (
	"project/config"
	"project/middleware"
	"project/models"
)

func GetDetailUsers(userId int) (interface{}, error) {
	user := models.User{}
	if e := config.DB.Find(&user, userId).Error; e != nil {
		return nil, e
	}
	return user, nil
}

func LoginUsers(user *models.User) (interface{}, error) {
	var err error
	if err = config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error; err != nil {
		return nil, err
	}
	userResp := models.UserResponse{Name: user.Name, Email: user.Email}
	userResp.Token, err = middleware.CreateToken(int(user.ID))
	if err != nil {
		return nil, err
	}
	if err := config.DB.Save(&user).Error; err != nil {
		return nil, err
	}
	return userResp, nil
}
