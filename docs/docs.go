// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api": {
            "get": {
                "tags": [
                    "App"
                ],
                "summary": "app route, get heathy status",
                "responses": {}
            }
        },
        "/api/v1/auth/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "login api",
                "parameters": [
                    {
                        "description": "Login inputs",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.LoginDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/core.Response-auth_TokensResponseType"
                        }
                    }
                }
            }
        },
        "/api/v1/auth/logout": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "logout user",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/core.SuccessResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/auth/password": {
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "change logged in user password",
                "parameters": [
                    {
                        "description": "Change password inputs",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.PasswordDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/core.SuccessResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/auth/refresh": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "refresh tokens api",
                "parameters": [
                    {
                        "description": "Refresh tokens inputs",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.RefreshDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/core.Response-auth_TokensResponseType"
                        }
                    }
                }
            }
        },
        "/api/v1/auth/register": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "register api",
                "parameters": [
                    {
                        "description": "Register inputs",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.RegisterDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/core.SuccessResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/auth/user": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "fetch logged in user info",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/core.Response-users_UserResponseType"
                        }
                    }
                }
            }
        },
        "/api/v1/communities": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Communities"
                ],
                "summary": "get list of communities",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/core.Response-communities_CommunitiesResponseType"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Communities"
                ],
                "summary": "create new community",
                "parameters": [
                    {
                        "description": "Create community inputs",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/communities.CreateDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/core.Response-communities_CommunityResponseType"
                        }
                    }
                }
            }
        },
        "/api/v1/communities/joined": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Communities"
                ],
                "summary": "get list of joined communities",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/core.Response-communities_CommunitiesResponseType"
                        }
                    }
                }
            }
        },
        "/api/v1/communities/own": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Communities"
                ],
                "summary": "get list of own communities",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/core.Response-communities_CommunitiesResponseType"
                        }
                    }
                }
            }
        },
        "/api/v1/communities/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Communities"
                ],
                "summary": "get community by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Community ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/core.Response-communities_CommunityResponseType"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Communities"
                ],
                "summary": "update community",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Community ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update community inputs",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/communities.UpdateDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/core.SuccessResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/communities/{id}/join": {
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Communities"
                ],
                "summary": "join to community",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Community ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/core.SuccessResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/communities/{id}/left": {
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Communities"
                ],
                "summary": "left from community",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Community ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/core.SuccessResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/events": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Events"
                ],
                "summary": "get list of events",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Community ID",
                        "name": "community_id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/core.Response-events_EventsResponseType"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Events"
                ],
                "summary": "create new event",
                "parameters": [
                    {
                        "description": "Create event inputs",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/events.CreateDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/core.Response-events_EventResponseType"
                        }
                    }
                }
            }
        },
        "/api/v1/events/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Events"
                ],
                "summary": "get event by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Event ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/core.Response-events_EventResponseType"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Events"
                ],
                "summary": "update event",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Event ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update event inputs",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/events.UpdateDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/core.SuccessResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/users": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "get list of users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/core.Response-users_UsersResponseType"
                        }
                    }
                }
            }
        },
        "/api/v1/users/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "get user by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/core.Response-users_UserResponseType"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "auth.LoginDto": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "maxLength": 90,
                    "minLength": 8
                }
            }
        },
        "auth.PasswordDto": {
            "type": "object",
            "required": [
                "current_password",
                "new_password"
            ],
            "properties": {
                "current_password": {
                    "type": "string",
                    "maxLength": 90,
                    "minLength": 8
                },
                "new_password": {
                    "type": "string",
                    "maxLength": 90,
                    "minLength": 8
                }
            }
        },
        "auth.RefreshDto": {
            "type": "object",
            "required": [
                "access_token",
                "refresh_token"
            ],
            "properties": {
                "access_token": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 30
                },
                "refresh_token": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 30
                }
            }
        },
        "auth.RegisterDto": {
            "type": "object",
            "required": [
                "email",
                "name",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string",
                    "maxLength": 190,
                    "minLength": 2
                },
                "password": {
                    "type": "string",
                    "maxLength": 90,
                    "minLength": 8
                }
            }
        },
        "auth.Tokens": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "expires_at": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "auth.TokensResponseType": {
            "type": "object",
            "properties": {
                "tokens": {
                    "$ref": "#/definitions/auth.Tokens"
                }
            }
        },
        "communities.CommunitiesResponseType": {
            "type": "object",
            "properties": {
                "communities": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/communities.Community"
                    }
                }
            }
        },
        "communities.Community": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "owner": {
                    "$ref": "#/definitions/users.UserShort"
                },
                "owner_id": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "users": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/users.UserShort"
                    }
                }
            }
        },
        "communities.CommunityResponseType": {
            "type": "object",
            "properties": {
                "community": {
                    "$ref": "#/definitions/communities.Community"
                }
            }
        },
        "communities.CreateDto": {
            "type": "object",
            "required": [
                "description",
                "title"
            ],
            "properties": {
                "description": {
                    "type": "string",
                    "maxLength": 250,
                    "minLength": 2
                },
                "title": {
                    "type": "string",
                    "maxLength": 190,
                    "minLength": 2
                }
            }
        },
        "communities.UpdateDto": {
            "type": "object",
            "required": [
                "description",
                "title"
            ],
            "properties": {
                "description": {
                    "type": "string",
                    "maxLength": 250,
                    "minLength": 2
                },
                "title": {
                    "type": "string",
                    "maxLength": 190,
                    "minLength": 2
                }
            }
        },
        "core.Response-auth_TokensResponseType": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/auth.TokensResponseType"
                },
                "errors": {
                    "type": "object"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "core.Response-communities_CommunitiesResponseType": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/communities.CommunitiesResponseType"
                },
                "errors": {
                    "type": "object"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "core.Response-communities_CommunityResponseType": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/communities.CommunityResponseType"
                },
                "errors": {
                    "type": "object"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "core.Response-events_EventResponseType": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/events.EventResponseType"
                },
                "errors": {
                    "type": "object"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "core.Response-events_EventsResponseType": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/events.EventsResponseType"
                },
                "errors": {
                    "type": "object"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "core.Response-users_UserResponseType": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/users.UserResponseType"
                },
                "errors": {
                    "type": "object"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "core.Response-users_UsersResponseType": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/users.UsersResponseType"
                },
                "errors": {
                    "type": "object"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "core.SuccessResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object"
                },
                "errors": {
                    "type": "object"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "events.CreateDto": {
            "type": "object",
            "required": [
                "community_id",
                "description",
                "duration",
                "start_at",
                "status",
                "title"
            ],
            "properties": {
                "community_id": {
                    "type": "string"
                },
                "description": {
                    "type": "string",
                    "maxLength": 250,
                    "minLength": 2
                },
                "duration": {
                    "type": "integer",
                    "maximum": 250,
                    "minimum": 1
                },
                "start_at": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "title": {
                    "type": "string",
                    "maxLength": 190,
                    "minLength": 2
                }
            }
        },
        "events.Event": {
            "type": "object",
            "properties": {
                "community_id": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "duration": {
                    "type": "integer"
                },
                "id": {
                    "type": "string"
                },
                "start_at": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "events.EventResponseType": {
            "type": "object",
            "properties": {
                "event": {
                    "$ref": "#/definitions/events.Event"
                }
            }
        },
        "events.EventsResponseType": {
            "type": "object",
            "properties": {
                "events": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/events.Event"
                    }
                }
            }
        },
        "events.UpdateDto": {
            "type": "object",
            "required": [
                "description",
                "duration",
                "start_at",
                "status",
                "title"
            ],
            "properties": {
                "description": {
                    "type": "string",
                    "maxLength": 250,
                    "minLength": 2
                },
                "duration": {
                    "type": "integer",
                    "maximum": 250,
                    "minimum": 1
                },
                "start_at": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "title": {
                    "type": "string",
                    "maxLength": 190,
                    "minLength": 2
                }
            }
        },
        "users.User": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "users.UserResponseType": {
            "type": "object",
            "properties": {
                "user": {
                    "$ref": "#/definitions/users.User"
                }
            }
        },
        "users.UserShort": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "users.UsersResponseType": {
            "type": "object",
            "properties": {
                "users": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/users.User"
                    }
                }
            }
        }
    },
    "securityDefinitions": {
        "Bearer": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
