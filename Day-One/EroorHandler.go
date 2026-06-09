package main

import (
	"errors"
	"fmt"
	"os"
	"io"
	"bufio"
)

func devide(a,b float64) (float64, error){
	if b == 0 {
		return 0, errors.New("dinomirator can not be zero!")
	}
	return a/b, nil
}

// Custom Errors
func validateAge(Age int) error{
	if Age < 18 {
		return errors.New("Age must be greater than or equal to 18.")
	}
	return nil
}

// mini project for login error handling
func login(username, password string) error{
	if username == ""{
		return errors.New("User Name is required!")
	}
	if password == "" {
		return errors.New("Password is required!")
	}
	if username!= "admin" || password != "abc123"{
		return errors.New("Invalid credintials!!")
	}
	return  nil
}

var invalidPassword = errors.New("Invalid Password!")
func checkPassword(password string) error{
	if password != "1234"{
		return fmt.Errorf("Password Error: %w",invalidPassword)
	}
	return nil
}

func main(){
	err:= checkPassword("1122")
	if err != nil{
		fmt.Println("Wrong Password")
	}

	if errors.Is(err, invalidPassword){
		fmt.Println("Wrong Password in errors.Is.")
	}

	result, error:= devide(10, 2)
	if error != nil{
		fmt.Println("Error: ", error)
		return
	}
	fmt.Println("Result: ", result)

	// file handling errors.
	file, err:= os.Open("test.txt")
	if err != nil{
		fmt.Println("Error: ", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	fmt.Println("File is opend success fully!!", file)

	content, errc:= io.ReadAll(file)
	if errc!= nil {
		fmt.Println("Error: ",errc)
		return
	}
	fmt.Println("File Content: ", content)

	// Custom Errors 
	erro:= validateAge(18)
	if erro!= nil {
		fmt.Println("Error: ", erro)
		return
	} else {
		fmt.Println("You are matured!!")
	}

	errl:=login("admin","abc123")
	if errl!= nil {
		fmt.Println("Login Error: ", errl)
	} else {
		fmt.Println("LogIn success full!!")
	}
}