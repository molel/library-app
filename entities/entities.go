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

type AuthorCreate struct {
	Name        string         `json:"name" binding:"required"`
	Surname     string         `json:"surname" binding:"required"`
	Description sql.NullString `json:"description"`
}

type AuthorUpdate struct {
	AuthorId    int            `json:"authorId"`
	Name        string         `json:"name"`
	Surname     string         `json:"surname"`
	Description sql.NullString `json:"description"`
}

type GenreCreate struct {
	GenreId int    `json:"genreId" binding:"required"`
	Name    string `json:"name" binding:"required"`
}

type GenreUpdate struct {
	Name string `json:"name" binding:"required"`
}

type BookCreate struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	GenreId     int    `json:"genreId" binding:"required"`
	AuthorId    int    `json:"authorId" binding:"required"`
}

type BookUpdate struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	GenreId     int    `json:"genreId"`
	AuthorId    int    `json:"authorId"`
}
