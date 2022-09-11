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

type UserCreate struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserGet struct {
	Id int `json:"id"`
	UserCreate
}

type AuthorCreate struct {
	Name        string `json:"name" binding:"required"`
	Surname     string `json:"surname" binding:"required"`
	Description string `json:"description"`
}

type AuthorUpdate struct {
	Name        *string `json:"name"`
	Surname     *string `json:"surname"`
	Description *string `json:"description"`
}

type AuthorGet struct {
	Id          int            `json:"id"`
	Name        string         `json:"name"`
	Surname     string         `json:"surname"`
	Description sql.NullString `json:"description"`
}

type GenreCreateAndGet struct {
	Id   int    `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type GenreUpdate struct {
	Name *string `json:"name"`
}

type BookCreate struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	GenreId     int    `json:"genreId" binding:"required"`
	AuthorId    int    `json:"authorId" binding:"required"`
}

type BookUpdate struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	GenreId     *int    `json:"genreId"`
	AuthorId    *int    `json:"authorId"`
}

type BookGet struct {
	Id          int            `json:"id"`
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
	GenreId     int            `json:"genreId"`
	AuthorId    int            `json:"authorId"`
}

type ListCreate struct {
	Title string `json:"title" binding:"required"`
}

type ListUpdate struct {
	Title *string `json:"title"`
}

type ListGet struct {
	Id     int    `json:"id"`
	UserId int    `json:"userId"`
	Title  string `json:"title"`
}

type ListItemCreate struct {
	BookId int    `json:"bookId" binding:"required"`
	Status string `json:"status" binding:"required"`
}

type ListItemGet struct {
	BookId int    `json:"bookId"`
	Status string `json:"status"`
}

type ListItemUpdate struct {
	Status *string `json:"status"`
}
