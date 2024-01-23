<!-- @format -->

# Gin Blog Project

Simple blog project with Golang and Gin framework.

## Technical Features

-   Gorm Orm
-   MVC pattern
-   Swagger api document
-   Custom JWT authentication
-   Custom commands
-   Middlewares
-   Faker data
-   Base response contract
-   Modular

## Clone and Run

```bash
$ git clone git@github.com:mst-ghi/gin-blog-project.git
$ cd ./gin-blog-project
$ cp .env.example .env #Update Database configuration
$ go mod download
$ go run . serve
```

## Commands

-   **serve** _Serve application_
-   **swag:init** _Generate swagger document_
-   **db:migrate** _Tables migration_
-   **db:seeder** _Insert seed data_

_Swagger document is available on :_ **/api/docs/index.html**

## Source structure

```bash
├── app
│   ├── articles
│   │   ├── articles.controller.go
│   │   ├── articles.response.go
│   │   ├── articles.routes.go
│   │   └── articles.service.go
│   ├── auth
│   │   ├── auth.controller.go
│   │   ├── auth.dto.go
│   │   ├── auth.responses.go
│   │   ├── auth.routes.go
│   │   └── auth.service.go
│   ├── users
│   │   ├── users.controller.go
│   │   ├── users.response.go
│   │   ├── users.routes.go
│   │   └── users.service.go
│   └── routing.go
├── core
│   ├── bootstrap
│   │   └── bootstrap.go
│   ├── cmd
│   │   ├── migrate.go
│   │   ├── root.go
│   │   ├── seed.go
│   │   ├── serve.go
│   │   └── swag.go
│   ├── config
│   │   └── config.go
│   ├── engine
│   │   └── engine.go
│   ├── middlewares
│   │   ├── cors.go
│   │   └── jwt.go
│   ├── swagger
│   │   └── swagger.go
│   ├── kernel.go
│   └── provider.go
├── database
│   ├── migrations
│   │   └── migrations.go
│   ├── models
│   │   ├── article.model.go
│   │   ├── token.model.go
│   │   └── user.model.go
│   ├── repositories
│   │   ├── article.repository.go
│   │   ├── token.repository.go
│   │   └── user.repository.go
│   ├── seeder
│   │   ├── article.seeder.go
│   │   ├── seeder.go
│   │   └── user.seeder.go
│   └── database.go
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── pkg
│   ├── handlers
│   │   └── error.go
│   ├── helpers
│   │   ├── converters.go
│   │   ├── crypto.go
│   │   └── strings.go
│   ├── messages
│   │   └── messages.go
│   └── validation
│       ├── messages.go
│       └── validation.go
├── .env.example
├── go.mod
├── go.sum
├── main.go
└── README.md
```

## Author

**[Mostafa Gholami](https://mst-ghi.github.io/)**

