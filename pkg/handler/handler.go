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
			authors.GET("/:author", h.getAuthorById)
			authors.PUT("/:author", h.updateAuthorById)
			authors.DELETE("/:author", h.deleteAuthorById)
		}
		genres := api.Group("/genres")
		{
			genres.POST("/", h.createGenre)
			genres.GET("/", h.getGenres)
			genres.GET("/:genre", h.getGenreById)
			genres.PUT("/:genre", h.updateGenreById)
			genres.DELETE("/:genre", h.deleteGenreById)
		}
		books := api.Group("/books")
		{
			books.POST("/", h.createBook)
			books.GET("/", h.getBooks)
			books.GET("/:book", h.getBookById)
			books.PUT("/:book", h.updateBookById)
			books.DELETE("/:book", h.deleteBookById)
		}
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getLists)
			//TODO getListById must return list itself and its items
			lists.GET("/:list", h.getListById)
			lists.PUT("/:list", h.updateListById)
			lists.DELETE("/:list", h.deleteListById)
			items := lists.Group(":list/items")
			{
				items.POST("/", h.createListItem)
				items.PUT("/:book", h.updateListItemById)
				items.DELETE("/:book", h.deleteListItemById)
			}
		}
	}
	return router
}
