// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Danyil-Mykola Obertan",
            "email": "danyilmykolaobertan@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/validate": {
            "post": {
                "description": "Validates the card number, expiration month, and expiration year, and returns whether the card is valid or not based on various checks including empty fields, format validation, and the Luhn algorithm.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CardValidation"
                ],
                "summary": "Validate card information",
                "parameters": [
                    {
                        "description": "Card Information",
                        "name": "requestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/validcard.RequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "valid\": true} \"Response when the card information is valid, error field may be null if there are no errors.",
                        "schema": {
                            "$ref": "#/definitions/validcard.ResponseBody"
                        }
                    },
                    "400": {
                        "description": "error\": {\"code\": 400, \"message\": \"invalid card number\"}, \"valid\": false} \"Response when there is an error in card validation.",
                        "schema": {
                            "$ref": "#/definitions/validcard.ResponseBody"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "validcard.Error": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "validcard.RequestBody": {
            "type": "object",
            "properties": {
                "cardNumber": {
                    "type": "string"
                },
                "expirationMonth": {
                    "type": "integer"
                },
                "expirationYear": {
                    "type": "integer"
                }
            }
        },
        "validcard.ResponseBody": {
            "type": "object",
            "properties": {
                "error": {
                    "$ref": "#/definitions/validcard.Error"
                },
                "valid": {
                    "type": "boolean"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "2.0",
	Host:        "localhost:8080",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "Credit Card Validator API",
	Description: "This is a simple API to validate credit card information such as card number, expiration month and expiration year",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}