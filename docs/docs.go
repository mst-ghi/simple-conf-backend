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
                        "type": "integer",
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
                    "type": "integer"
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
