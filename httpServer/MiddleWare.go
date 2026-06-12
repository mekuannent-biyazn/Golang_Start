package main

import (
	"net/http"
	"fmt"
	"time"
)

func loger(
	next http.HandlerFunc,
)http.HandlerFunc{

	return func(
		w http.ResponseWriter,
		r *http.Request,
	){
		fmt.Fprintln(w,"Request is received! ")

		next(w, r)
	}
}

func login(
	next http.HandlerFunc,
)http.HandlerFunc{
	return func(
		w http.ResponseWriter, r *http.Request,
	){
		fmt.Fprintln(w, "You are authenticated user! ")
		next(w, r)
	}
}

func Authorize(
	w http.ResponseWriter, 
	r *http.Request,
){
	fmt.Fprintln(w,"You are authorized to this privelage!!")
}

func home(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w,"Well Come To Home")
}

func timer(
	next http.HandlerFunc,
)http.HandlerFunc{
	return func(
		w http.ResponseWriter,
		r *http.Request,
	){
		start:= time.Now()

		next(w,r)
		fmt.Println("Excution Takes: ", time.Since(start),
		)
	}
}

func auth(
	next http.HandlerFunc,
)http.HandlerFunc{
	return func(
		w http.ResponseWriter,
		r *http.Request,
	){
		token:= r.Header.Get(
			"Authorization",
		)
    //    authHeader := r.Header.Get("Authorization")

	// 	if len(authHeader) < 7 || authHeader[:7] != "Bearer " {
    //         http.Error(w, "Unauthorized", http.StatusUnauthorized)
    //         return
    //     }
        
    //     token := authHeader[7:]

		if token != "secret123"{
			http.Error(
				w,
				"Unauthorized",
				http.StatusUnauthorized,
			)
			return
		}
		next(w, r)
	}
}

func main(){
	http.HandleFunc("/", loger(home))
	http.HandleFunc("/auth", loger(login(Authorize)))
	http.HandleFunc("/authentication", auth(home))
	http.HandleFunc("/time", timer(home))
	fmt.Println("Srver is running on http://localhost:7070")
	http.ListenAndServe(":7070",nil)
}