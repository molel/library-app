package entities

import (
	"database/sql"
	"net/http"
	"time"
)

type HTTPServerConfigs struct {
	Addr           string
	Handler        http.Handler
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	MaxHeaderBytes int
}

type DatabaseConfigs struct {
	Host     string
	Port     int
	User     string
	Name     string
	Password string
	SSLMode  string
}

type UserSignUp struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type User struct {
	UserId int `json:"userId"`
	UserSignUp
}

type Author struct {
	AuthorId    int            `json:"authorId"`
	Name        string         `json:"name" binding:"required"`
	Surname     string         `json:"surname" binding:"required"`
	Description sql.NullString `json:"description"`
}
