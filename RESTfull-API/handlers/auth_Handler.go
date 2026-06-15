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
	type UserResponse struct{
		Id int `json: "id"`
		Name string `json: "name"`
		Email string `json: "email"`
	}

	if r.Method!= http.MethodPost{
		http.Error(
			w,
			"Method not allowed!",
			http.StatusMethodNotAllowed,
		)
		return
	}

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

	for _,u:= range users{
		if u.Email== user.Email{
			http.Error(
				w,
				"Email is already exist!",
				http.StatusBadRequest,
			)
			return
		}
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

	response:= UserResponse{
		Id: user.Id,
		Name: user.Name,
		Email: user.Email,
	}
	json.NewEncoder(w).Encode(&response)
}

func Login(
	w http.ResponseWriter,
	r *http.Request,
){
	var loginData struct{
		Email string `json: "email"`
		Password string `json: "password"`
	}

	if r.Method!= http.MethodPost{
		http.Error(
			w,
			"Method not allowed!",
			http.StatusMethodNotAllowed,
		)
		return
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