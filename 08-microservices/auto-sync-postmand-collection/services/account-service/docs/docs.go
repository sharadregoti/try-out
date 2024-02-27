// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
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
        "/accounts": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Get accounts of a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "GetAccounts",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Id to start after",
                        "name": "starting_after",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Id to start before",
                        "name": "ending_before",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "The number of records to return. Default: 20",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "optional userId for testing",
                        "name": "user_id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/http.AccountsResponse"
                        }
                    }
                }
            }
        },
        "/accounts/{accountId}/transactions": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Get accounts of a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "GetTransanctions",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Account ID",
                        "name": "accountId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Id to start after",
                        "name": "starting_after",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Id to start before",
                        "name": "ending_before",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "The number of records to return. Default: 20",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "http.Account": {
            "type": "object",
            "properties": {
                "balance_decimal_formatted": {
                    "type": "string"
                },
                "balance_formatted": {
                    "type": "string"
                },
                "balance_main_formatted": {
                    "type": "string"
                },
                "balance_raw": {
                    "type": "integer"
                },
                "country_flag": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "currency_symbol": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "kyc": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "number": {
                    "type": "string"
                },
                "provider_id": {
                    "type": "string"
                },
                "refreshed_at": {
                    "type": "string"
                },
                "service_status": {
                    "type": "integer"
                },
                "symbol_left": {
                    "type": "boolean"
                },
                "symbol_right": {
                    "type": "boolean"
                },
                "type": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "http.AccountsResponse": {
            "type": "object",
            "properties": {
                "accounts": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/http.Account"
                    }
                },
                "has_more": {
                    "type": "boolean"
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
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            " B2C API",
	Description:      "This is  server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
