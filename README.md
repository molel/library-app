# library-app

The project is a restful api service designed for a library management system

## Installing / Getting started

```shell
cd path_to_project
go build ./cmd/main.go
main
```

## Developing

### Built With

The project uses **Gin Web Framework**, **PostgreSQL Server**, **pq** as postgres driver, **sqlx** as an extension of
sql package from golang standard library, **jwt-go** as golang implementation of JSON Web Tokens, **GoDotEnv** and **Viper**

### Setting up Dev

Here's a brief intro about what a developer must do in order to start developing
the project further:

```shell
git clone https://github.com/molel/library-app
cd path_to_project
go mod download
```

## Configuration

Before the start it would be good practice to set up configurations in config file. At this moment server and database
port and other minor settings can be configured.

## Api Reference

In development
