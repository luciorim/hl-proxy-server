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
        "/proxy": {
            "post": {
                "description": "allows make http request to another website",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "proxy-controller"
                ],
                "summary": "process requested url",
                "parameters": [
                    {
                        "description": "request data",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.ProxyRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.ProxyResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "502": {
                        "description": "Bad Gateway",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/proxy/{id}": {
            "get": {
                "description": "return request and response by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "proxy-controller"
                ],
                "summary": "get proxy history by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "request id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ProxyResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.ProxyRequest": {
            "type": "object",
            "properties": {
                "headers": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "method": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "dto.ProxyResponse": {
            "type": "object",
            "properties": {
                "headers": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "array",
                        "items": {
                            "type": "string"
                        }
                    }
                },
                "id": {
                    "type": "string"
                },
                "length": {
                    "type": "integer"
                },
                "status": {
                    "type": "integer"
                },
                "url": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8181",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Proxy Server",
	Description:      "Server allows do http requests to another websites",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
