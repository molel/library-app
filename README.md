# library-app
Сервер для обработки данных в системе хранения книг библиотеки

***

Для запуска используется следующие команды:
```
> go build path_to_project/cmd/main.go

> main
```

***

## API (в плане):

`POST auth/sign-up`
Регистрация пользователей в системе

`POST auth/sign-in`
Авторизация пользователей в системе

`GET api/books/`
Получение всех сущностей книг

`GET api/books/id`
Получение сущности книги по ее id

`GET api/books/author-id`
Получение сущности книги по id сущности ее автора

`POST api/books/`
Добавление сущности книги

`PUT api/books/id`
Изменение сущности книги по ее id

`DELETE api/books/id`
Удаление сущности книги по ее id

`GET api/authors/`
Получение всех сущностей авторов

`GET api/authors/id`
Получение сущности автора по ее id

`POST api/authors/`
Добавление сущности автора

`PUT api/authors/id`
Изменение сущности автора по ее id

`DELETE api/authors/id`
Удаление сущности автора по ее id

