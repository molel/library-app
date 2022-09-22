package repository

import (
	"fmt"
	"library-app/entities"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewPostgresqlDB(configs entities.DatabaseConfigs) *sqlx.DB {
	db, err := sqlx.Connect("postgres",
		fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
			configs.Host, configs.Port, configs.User, configs.Name, configs.Password, configs.SSLMode))
	if err != nil {
		log.Fatalf("error occurred while connecting to database: %s\n", err.Error())
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("error occurred while pinging the database: %s\n", err.Error())
	}
	return db
}
