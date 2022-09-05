package handler

import (
	"github.com/gin-gonic/gin"
	"library-app/pkg/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Init() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.SignUp)
		auth.POST("/sign-in", h.SignIn)
	}
	api := router.Group("/api", h.UserIdentification)
	{
		authors := api.Group("/authors")
		{
			authors.POST("/", h.createAuthor)
			authors.GET("/", h.getAuthors)
			authors.GET("/:id", h.getAuthorById)
			authors.PUT("/:id", h.updateAuthorById)
			authors.DELETE("/:id", h.deleteAuthorById)
		}
		genres := api.Group("/genres")
		{
			genres.POST("/", h.createGenre)
			genres.GET("/", h.getGenres)
			genres.GET("/:id", h.getGenreById)
			genres.PUT("/:id", h.UpdateGenreById)
			genres.DELETE("/:id", h.deleteGenreById)
		}
		books := api.Group("/books")
		{
			books.POST("/", h.createBook)
			books.GET("/", h.getBooks)
			books.GET("/:id", h.getBookById)
			books.PUT("/:id", h.updateBookById)
			books.DELETE("/:id", h.deleteBookById)
		}
		// TODO implement handlers for users' book list
	}
	return router
}
