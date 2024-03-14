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
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/orders": {
            "post": {
                "description": "Creating Order with given request",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Orders"
                ],
                "summary": "Creating Order",
                "parameters": [
                    {
                        "type": "string",
                        "description": "code of Order",
                        "name": "x-correlationid",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Request of Creating Order Object",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreateOrderCommand"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/orders/{id}": {
            "get": {
                "description": "Getting Order by Id in detail",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Orders"
                ],
                "summary": "Getting Order by Id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id of Order",
                        "name": "x-correlationid",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "id of Order",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.CreateOrderCommand": {
            "description": "Request about creating Order",
            "type": "object",
            "required": [
                "cargoId",
                "lineItems",
                "shipmentNumber"
            ],
            "properties": {
                "cargoId": {
                    "description": "cargo id of Order",
                    "type": "integer"
                },
                "lineItems": {
                    "description": "cargo id of Order",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.CreateOrderLineItemCommand"
                    }
                },
                "shipmentNumber": {
                    "description": "shipment no of Order",
                    "type": "integer"
                }
            }
        },
        "model.CreateOrderLineItemCommand": {
            "type": "object",
            "required": [
                "productId",
                "sellerId"
            ],
            "properties": {
                "productId": {
                    "description": "product id of Order line items",
                    "type": "integer"
                },
                "sellerId": {
                    "description": "product id of Order line items",
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Order Api",
	Description:      "This is an Order Api just for young people",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
