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
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/organizations/": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.OrganizationDB"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.OrganizationDB": {
            "type": "object",
            "properties": {
                "about": {
                    "type": "string"
                },
                "color": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "employees_amount": {
                    "type": "integer"
                },
                "full_name": {
                    "$ref": "#/definitions/sql.NullString"
                },
                "gr_points": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "logo": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "organization_inn": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "revenue": {
                    "type": "integer"
                },
                "status": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "sql.NullString": {
            "type": "object",
            "properties": {
                "string": {
                    "type": "string"
                },
                "valid": {
                    "description": "Valid is true if String is not NULL",
                    "type": "boolean"
                }
            }
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
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
