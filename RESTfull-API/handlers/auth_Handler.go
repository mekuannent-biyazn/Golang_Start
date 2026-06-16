package handlers


import (
	"encoding/json"
	"net/http"
	"fmt"
	"strconv"

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
		Id int `json:"id"`
		Name string `json:"name"`
		Email string `json:"email"`
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
		Email string `json:"email"`
		Password string `json:"password"`
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
			"token": token,
		},
	)
}

func Profile(
	w http.ResponseWriter,
	r *http.Request,
){
	userID:= r.Context().Value("userID")

	fmt.Println("Profile UserID:", userID)

	json.NewEncoder(w).
		Encode(
			map[string]interface{}{
				"message":
					"Welcome",
				"user_id":
					userID,
			},
		)
}

func UpdateProfile(
	w http.ResponseWriter,
	r *http.Request,
){
	var user models.User
	var updateData struct{
		Name string`json:"name"`
		Email string`json:"email"`
		Password string`json:"passsword"`
	}

	if r.Method!= http.MethodPut{
		http.Error(
			w,
			"Method not allowed!!",
			http.StatusMethodNotAllowed,
		)
		return
	}

	json.NewDecoder(
		r.Body,
	).Decode(&updateData)

	user.Name = updateData.Name
	user.Email = updateData.Email

	hash, err := utils.HashPassword(
		updateData.Password,
	)

	if err!= nil{
		http.Error(
			w,
			"Password Error!",
			http.StatusInternalServerError,
		)
	}
	user.Password = hash

	users= append(users, user)

	json.NewEncoder(w).Encode(
		map[string]interface{}{
				"message":
					"Your Profile data is changed successfully!!",
				"New Name":
					user.Name,
			},
	)
}

func DeleteUser(
	w http.ResponseWriter,
	r *http.Request,
){
	ID:= r.URL.Query().Get("id")

	if r.Method!= http.MethodDelete{
		http.Error(
			w,
			"Mothod not allowed!",
			http.StatusMethodNotAllowed,
		)
		return
	}

	if ID== ""{
		http.Error(
			w,
			"User Id is requered.",
			http.StatusBadRequest,
		)
		return
	}

	userId, err := strconv.Atoi(ID)

	if err != nil {
		http.Error(
			w,
			"Invalid ID",
			http.StatusBadRequest,
		)
		return
	}

	userIndex:= -1

	for i, user := range users{
		if user.Id == userId{
			userIndex = i
			break
		}
	}

	if userIndex == -1{
		http.Error(
			w,
			"User Not Found!",
			http.StatusNotFound,
		)
		return
	}

	users = append(
		users[:userIndex],
		users[userIndex+1:]...,
	)

	json.NewEncoder(w).Encode(
		map[string]string{
			"message": "User Deleted Successfully",
		},
	)

}

func GetAllUsers(
	w http.ResponseWriter,
	r *http.Request,
){
	json.NewEncoder(w).Encode(&users)
}