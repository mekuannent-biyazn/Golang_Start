package main

import (
	"errors"
	"fmt"
	"os"
	"io"
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

func main(){
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

	content, errc:= io.ReadAll(file)
	if errc!= nil {
		fmt.Println("Error: ",errc)
		return
	}
	fmt.Println("File Content: ", content)

	defer file.Close()

	fmt.Println("File is opend success fully!!", file)

	// Custom Errors 
	erro:= validateAge(18)
	if erro!= nil {
		fmt.Println("Error: ", erro)
		return
	} else {
		fmt.Println("You are matured!!")
	}
}