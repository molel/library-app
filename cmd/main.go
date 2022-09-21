package main

import (
	server "library-app"
	"library-app/config"
	"library-app/entities"
	"library-app/pkg/handler"
	"library-app/pkg/repository"
	"library-app/pkg/service"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

// @title Library app API
// @version 1.0
// @description RESTful API service for library management system

// @host localhost:8000
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	if err := config.InitConfigs(); err != nil {
		log.Fatalf("error occured while reading in configs: %s\n", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error occured while loading env variables: %s\n", err.Error())
	}

	databaseConfigs := entities.DatabaseConfigs{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetInt("db.port"),
		User:     viper.GetString("db.user"),
		Name:     viper.GetString("db.name"),
		Password: os.Getenv("DB_PASSWORD"),
		SSLMode:  viper.GetString("db.sslMode"),
	}

	db := repository.NewPostgresqlDB(databaseConfigs)

	repositories := repository.NewRepository(db)

	services := service.NewService(repositories)

	handlers := handler.NewHandler(services)

	httpServerConfigs := &entities.HTTPServerConfigs{
		Addr:           "localhost:" + viper.GetString("server.port"),
		Handler:        handlers.Init(),
		ReadTimeout:    time.Duration(viper.GetInt("server.readTimeout")) * time.Second,
		WriteTimeout:   time.Duration(viper.GetInt("server.writeTimeout")) * time.Second,
		MaxHeaderBytes: viper.GetInt("server.maxHeaderKilobytes") << 10,
	}

	if err := new(server.Server).Run(httpServerConfigs); err != nil {
		log.Fatalf("error occured while running the server: %s\n", err.Error())
	}
}
