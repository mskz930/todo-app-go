package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"todo-app/model/dto"
	"todo-app/service"

	_ "github.com/go-sql-driver/mysql"
)

var loginService service.LoginService

func main() {
	db, err := sql.Open("mysql", "root@(127.0.0.1:3306)/test")
	if err != nil {
		panic(err)
	}
	loginService = service.NewLoginService(*db)

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/", index)
	http.HandleFunc("/login", login)
	server.ListenAndServe()
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		response := &dto.Response[any]{Status: 200, Message: "Ok"}
		json.NewEncoder(w).Encode(response)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "POST" {
		return
	}
	loginRequest := dto.LoginRequest{}
	json.NewDecoder(r.Body).Decode(&loginRequest)

	user, err := loginService.Login(loginRequest)
	var response *dto.Response[*dto.User]
	if user != nil || err != nil {
		response = &dto.Response[*dto.User]{Status: 200, Message: "login success", Data: user}
	} else {
		response = &dto.Response[*dto.User]{Status: 200, Message: "login failed"}
	}
	json.NewEncoder(w).Encode(response)
}
