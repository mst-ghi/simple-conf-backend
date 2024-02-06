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
$ git clone git@github.com:mst-ghi/video-conf-backend.git
$ cd ./video-conf-backend
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
│   ├── auth
│   │   ├── auth.controller.go
│   │   ├── auth.dto.go
│   │   ├── auth.responses.go
│   │   ├── auth.routes.go
│   │   └── auth.service.go
│   ├── communities
│   │   ├── communities.controller.go
│   │   ├── communities.dto.go
│   │   ├── communities.response.go
│   │   ├── communities.routes.go
│   │   └── communities.service.go
│   ├── events
│   │   ├── events.controller.go
│   │   ├── events.dto.go
│   │   ├── events.response.go
│   │   ├── events.routes.go
│   │   └── events.service.go
│   ├── rooms
│   │   ├── events.controller.go
│   │   ├── events.routes.go
│   │   ├── rooms.dto.go
│   │   ├── rooms.response.go
│   │   └── rooms.service.go
│   ├── users
│   │   ├── users.controller.go
│   │   ├── users.response.go
│   │   ├── users.routes.go
│   │   └── users.service.go
│   ├── gateway.go
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
│   ├── socket
│   │   ├── constants.go
│   │   ├── emiter.go
│   │   ├── generator.go
│   │   ├── handlers.go
│   │   ├── middleware.go
│   │   └── socket.go
│   ├── swagger
│   │   └── swagger.go
│   ├── kernel.go
│   └── provider.go
├── database
│   ├── migrations
│   │   └── migrations.go
│   ├── models
│   │   ├── community.model.go
│   │   ├── event.model.go
│   │   ├── room.model.go
│   │   ├── token.model.go
│   │   └── user.model.go
│   ├── repositories
│   │   ├── community.repository.go
│   │   ├── event.repository.go
│   │   ├── room.repository.go
│   │   ├── token.repository.go
│   │   └── user.repository.go
│   ├── seeder
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
├── go.mod
├── go.sum
├── main.go
└── README.md
```

## Author

**[Mostafa Gholami](https://mst-ghi.github.io/)**

