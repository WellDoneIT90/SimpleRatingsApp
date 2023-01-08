// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "email": "your@mail.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/rating": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create a new rating.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Rating"
                ],
                "summary": "create a new rating",
                "parameters": [
                    {
                        "description": "Company",
                        "name": "company",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "Author",
                        "name": "author",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "AuthorRole",
                        "name": "author_role",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "CompanyRating",
                        "name": "company_rating",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "integer"
                        }
                    },
                    {
                        "description": "Description",
                        "name": "description",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Ratings"
                        }
                    }
                }
            }
        },
        "/v1/rating/": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Delete rating by given ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Rating"
                ],
                "summary": "delete rating by given ID",
                "parameters": [
                    {
                        "description": "Rating ID",
                        "name": "id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/rating/{id}": {
            "get": {
                "description": "Get rating by given ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Rating"
                ],
                "summary": "get rating by given ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Rating ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Ratings"
                        }
                    }
                }
            }
        },
        "/v1/ratings": {
            "get": {
                "description": "Get all existing ratings.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Ratings"
                ],
                "summary": "get all existing ratings",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Ratings"
                            }
                        }
                    }
                }
            }
        },
        "/v1/token/new": {
            "get": {
                "description": "Create a new access token.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Token"
                ],
                "summary": "create a new access token",
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Ratings": {
            "type": "object",
            "required": [
                "author",
                "authorrole",
                "company",
                "description",
                "id",
                "rating"
            ],
            "properties": {
                "author": {
                    "type": "string",
                    "maxLength": 255
                },
                "authorrole": {
                    "type": "string",
                    "maxLength": 255
                },
                "company": {
                    "type": "string",
                    "maxLength": 255
                },
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string",
                    "maxLength": 255
                },
                "id": {
                    "type": "string"
                },
                "rating": {
                    "type": "integer",
                    "maximum": 5,
                    "minimum": 1
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "API",
	Description:      "This is an auto-generated API Docs.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
