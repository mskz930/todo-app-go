package service

import (
	"database/sql"
	"errors"
	"log"
	"todo-app/model/dto"
)

type LoginService interface {
	Login(dto.LoginRequest) (*dto.User, error)
}

type loginService struct {
	db *sql.DB
}

func NewLoginService(db sql.DB) LoginService {
	return &loginService{db: &db}
}

func (ls *loginService) Login(loginRequest dto.LoginRequest) (*dto.User, error) {
	sql := "select id, name, password from users where email = ?"
	var (
		id       int
		name     string
		password string
	)

	err := ls.db.QueryRow(sql, loginRequest.Email).Scan(&id, &name, &password)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	// パスワードが誤りの場合
	if loginRequest.Password != password {
		return nil, errors.New("password is invalid")
	}
	// 認証時
	return &dto.User{Id: id, Name: name, Email: loginRequest.Email}, nil
}
