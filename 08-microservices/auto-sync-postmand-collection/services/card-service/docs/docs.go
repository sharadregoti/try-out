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
        "/cards": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Get cards of a user",
                "summary": "Getcards",
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
                            "$ref": "#/definitions/http.CardsResponse"
                        }
                    }
                }
            }
        },
        "/cards/{cardId}/block": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "This is the first line",
                "summary": "BlockCard",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Card ID",
                        "name": "cardId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "query params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.BlockCardRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/http.EmptyResponse"
                        }
                    }
                }
            }
        },
        "/cards/{cardId}/cvv-toogle": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "This is the first line",
                "summary": "GetCardCVV",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Card ID",
                        "name": "cardId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/http.CardCVVResponse"
                        }
                    }
                }
            }
        },
        "/cards/{cardId}/features/{featuredId}/limit": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "This is the first line",
                "summary": "SetFeatureLimit",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Card ID",
                        "name": "cardId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Card ID",
                        "name": "featuredId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "query params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.SetFeatureLimitRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/http.EmptyResponse"
                        }
                    }
                }
            }
        },
        "/cards/{cardId}/features/{featuredId}/status": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "This is the first line",
                "summary": "SetFeatureStatus",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Card ID",
                        "name": "cardId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Card ID",
                        "name": "featuredId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "query params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.SetFeatureEnableRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/http.EmptyResponse"
                        }
                    }
                }
            }
        },
        "/cards/{cardId}/pin": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "This is the first line",
                "summary": "SetCardPin",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Card ID",
                        "name": "cardId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "query params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.SetPinRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/http.EmptyResponse"
                        }
                    }
                }
            }
        },
        "/cards/{cardId}/rewards": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "This is the first line",
                "summary": "GetCardRewards",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Id to start after",
                        "name": "start_date",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Id to start before",
                        "name": "end_date",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "card type",
                        "name": "card_type",
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
                        "description": "card ID",
                        "name": "cardId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/services.GetRewardPointsResponse"
                        }
                    }
                }
            }
        },
        "/cards/{cardId}/transactions": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Get cards of a user",
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
                        "description": "Card ID",
                        "name": "cardId",
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
        "http.BlockCardRequest": {
            "type": "object",
            "required": [
                "isBlocked"
            ],
            "properties": {
                "isBlocked": {
                    "type": "boolean"
                }
            }
        },
        "http.Card": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "current_balance": {
                    "type": "integer"
                },
                "current_balance_formatted": {
                    "type": "string"
                },
                "expiry": {
                    "type": "string"
                },
                "features": {
                    "$ref": "#/definitions/http.CardFeatures"
                },
                "id": {
                    "type": "string"
                },
                "is_addon_enabled": {
                    "type": "boolean"
                },
                "is_blocked": {
                    "type": "boolean"
                },
                "is_reward_enabled": {
                    "type": "boolean"
                },
                "limit": {
                    "type": "integer"
                },
                "limit_formatted": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "number": {
                    "type": "string"
                },
                "processor": {
                    "type": "string"
                },
                "total_reward_points": {
                    "type": "integer"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "http.CardCVVResponse": {
            "type": "object",
            "properties": {
                "cvv": {
                    "type": "integer"
                }
            }
        },
        "http.CardFeature": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "is_enabled": {
                    "type": "boolean"
                },
                "limit": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "http.CardFeatures": {
            "type": "object",
            "properties": {
                "domestic": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/http.CardFeature"
                    }
                },
                "international": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/http.CardFeature"
                    }
                }
            }
        },
        "http.CardsResponse": {
            "type": "object",
            "required": [
                "cards"
            ],
            "properties": {
                "cards": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/http.Card"
                    }
                },
                "has_more": {
                    "type": "boolean"
                }
            }
        },
        "http.EmptyResponse": {
            "type": "object"
        },
        "http.SetFeatureEnableRequest": {
            "type": "object",
            "required": [
                "isEnabled"
            ],
            "properties": {
                "isEnabled": {
                    "type": "boolean"
                }
            }
        },
        "http.SetFeatureLimitRequest": {
            "type": "object",
            "required": [
                "limit"
            ],
            "properties": {
                "limit": {
                    "type": "integer"
                }
            }
        },
        "http.SetPinRequest": {
            "type": "object",
            "required": [
                "pin"
            ],
            "properties": {
                "pin": {
                    "type": "integer"
                }
            }
        },
        "services.GetRewardPointsResponse": {
            "type": "object",
            "properties": {
                "has_more": {
                    "type": "boolean"
                },
                "reward_points": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/services.RewardPoint"
                    }
                }
            }
        },
        "services.RewardPoint": {
            "type": "object",
            "properties": {
                "card_id": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "long_description": {
                    "type": "string"
                },
                "points": {
                    "type": "integer"
                },
                "short_description": {
                    "type": "string"
                },
                "transaction_id": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
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
