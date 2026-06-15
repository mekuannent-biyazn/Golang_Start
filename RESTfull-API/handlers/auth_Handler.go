package handlers


import (
	"encoding/json"
	"net/http"

	"go-auth-api/models"
	"go-auth-api/utils"
)

var users []models.User

func Register(
	w http.ResponseWriter, 
	r *http.Request,
){
	var user models.User

	err:= json.NewDecoder(
		r.Body,
	).Decode(&user)

	if err!= nil {
		http.Error(
			w,
			"Invalid Input",
			http.StatusBadRequest,
		)
		return
	}

	if user.Name=="" || user.Email=="" || user.Password==""{
		http.Error(
			w,
			"Please fill the requered fields!",
			http.StatusBadRequest,
		)
		return
	}

	hashPassword, err:= utils.HashPassword(
		user.Password,
	)
	if err!= nil {
		http.Error(
			w,
			"Password Hashing Error!",
			http.StatusInternalServerError,
		)
		return
	}

	user.Password = hashPassword
	user.Id = len(users)+1

	users= append(users, user,)

	w.WriteHeader(
		http.StatusCreated,
	)
	json.NewEncoder(w).Encode(&user)
}

func Login(
	w http.ResponseWriter,
	r *http.Request,
){
	var loginData struct{
		Email string `json: "email"`
		Password string `json: "password"`
	}

	json.NewDecoder(
		r.Body,
	).Decode(&loginData)

	var foundUser models.User
	var userFound bool

	for _, user := range users{
		if user.Email == loginData.Email{
			foundUser = user
			userFound = true

			break
		}
	}

	if !userFound{
		http.Error(
			w,
			"User Not Found!",
			http.StatusNotFound,
		)
		return
	}

	err := utils.CheckPassword(
		foundUser.Password,
		loginData.Password,
	)

	if err!= nil {
		http.Error(
			w,
			"Invalid Credintial!!",
			http.StatusUnauthorized,
		)
		return
	}

	token, err:= utils.GenerateToken(
		foundUser.Id,
	)

	if err!= nil {
		http.Error(
			w,
			"Token Error!!",
			http.StatusInternalServerError,
		)
		return
	}

	json.NewEncoder(w).Encode(
		map[string]string{
			"token: ": token,
		},
	)
}