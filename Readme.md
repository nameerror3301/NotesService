<p align="center">
      <img src="./image/Логотип%20для%20проекта.png" width="725">
</p>

<p align="center">
   <img src="https://img.shields.io/badge/Golang-version%201.19-blue" alt="Golang">
   <img src="https://img.shields.io/badge/NotesService-version%201.0-blue" alt="RenamePhotoInDir">
</p>

# Notes Service

## Описание

Данный проект был реализован с целью совершенствования навыков в написании REST API. 

Использовались только библиотеки языка Golang.

В корневой директории есть файл под названием **"Terms of Reference.md"** в нем описанно ТЗ проекта.

## Сборка 

Для сборки проекта вы можете использовать Docker для этого в корневой директории пропишите команду **"docker-compose up --build"**.

Или же вы можете запустить это на прямую, для этого в корневой директории пропишите команду **"go run ./app/cmd/*.go"**.

## Структура проекта

```
.
├── app
│   └── cmd
│       └── main.go
├── internal
│   ├── app
│   │   └── run
│   │       └── run.go
│   ├── middleware
│   │   └── user.go
│   ├── models
│   │   ├── notes.go
│   │   └── users.go
│   └── routes
│       ├── auth.go
│       ├── notes.go
│       └── route.go
├── docker-compose.yaml
├── Dockerfile
├── go.mod
├── go.sum
├── Readme.md
├── TermsOfReference.md
└── tools
    └── crypt.go
```

## API точки доступа

Ниже будут перечисленны точки доступа к API.

Везде где требуется передать данные методами **POST** **PUT** тип передаваемых данных это **application/json** в ином случае вы получите ошибку.


**URL** - http://localhost:8080

#### **/api/v1**
* `GET` : Получение стартовой страницы

#### **/api/v1/signUp**
* `POST` : Регистрация средствами BasicAuth.

#### **/api/v1/notes**
* `GET` : Получить все заметки
* `AUTH` : Требуется

#### **/api/v1/notes?id={number}**
* `GET` : Получение заметки по ее идентификатору
* `AUTH` : Требуется

#### **/api/v1/notes?sort={ASC || DESC}**
* `GET` : Получение отсортированных значений (Сортировка по идентификатору)
* `AUTH` : Требуется

#### **/api/v1/notes/create**
* `POST` : Создание заметки
* `AUTH` : Требуется
```
    {
        "notes_name": "Note #1", 
        "notes_content": "Hello world!",
        "ttl_date_to_delete":"2022/12/04 12:27:44" 
    }
```

#### **/api/v1/notes/upload**
* `PUT` : Обновление заметки (Можно обновить имя или контент, или все сразу)
* `AUTH` : Требуется 
```
    {
        "id": 1, // Айди заметки которую хотим обновить
        "notes_name": "Note #2", 
        "notes_content": "Bye Bye World!"
    }
```

#### **/api/v1/notes/delete?id={number}**
* `DELETE` : Удаление заметки по ее идентификатору
* `AUTH` : Требуется
