migration_directory=./schema
database_login=postgres
database_password=qwerty123
database_port=5432
database_name=postgres

start: build up

stop: down rm

build:
	docker-compose build

up:
	docker-compose up -d

rm:
	docker-compose rm

down:
	docker-compose down

migrate:
	migrate -path $(migration_directory) -database "postgres://$(database_login):$(database_password)@localhost:$(database_port)/${database_name}?sslmode=disable" up >> ./logs/migrations.txt 2>&1
