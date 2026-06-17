package service

import (
	"errors"
	"golangpkg/models"
	"golangpkg/repositories"
	"golangpkg/utils"
)

func RegisterUser(user models.User) error {
	hashedPass, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPass
	return repositories.CreateUser(user)
}

func LoginUser(email, password string) (string, error) {
	user, err := repositories.GetUserByEmail(email)
	if err != nil {
		return "", err
	}

	err = utils.CheckPassword(user.Password, password)
	if err != nil {
		return "", errors.New("Invalid Credential!!")
	}

	token, err := utils.GenerateToken(user.Id)
	if err != nil {
		return "", err
	}
	return token, nil
}

func GetProfile(userID int) (models.User, error) {
	return repositories.GetUserByID(userID)
}