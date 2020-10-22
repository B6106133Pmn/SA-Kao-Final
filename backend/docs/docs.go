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
        "/degrees": {
            "get": {
                "description": "list degree entities",
                "produces": [
                    "application/json"
                ],
                "summary": "List degree entities",
                "operationId": "list-degree",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Offset",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/ent.Degree"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            },
            "post": {
                "description": "Create degree",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create degree",
                "operationId": "create-degree",
                "parameters": [
                    {
                        "description": "Degree entity",
                        "name": "degree",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/ent.Degree"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ent.Degree"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/degrees/{id}": {
            "get": {
                "description": "get degree by ID",
                "produces": [
                    "application/json"
                ],
                "summary": "Get a degree entity by ID",
                "operationId": "get-degree",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Degree ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ent.Degree"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/departments": {
            "get": {
                "description": "list department entities",
                "produces": [
                    "application/json"
                ],
                "summary": "List department entities",
                "operationId": "list-department",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Offset",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/ent.Department"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            },
            "post": {
                "description": "Create department",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create department",
                "operationId": "create-department",
                "parameters": [
                    {
                        "description": "Department entity",
                        "name": "department",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/ent.Department"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ent.Department"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/departments/{id}": {
            "get": {
                "description": "get doctor by ID",
                "produces": [
                    "application/json"
                ],
                "summary": "Get a doctor entity by ID",
                "operationId": "get-doctor",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Doctor ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ent.Doctor"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            },
            "put": {
                "description": "update department by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Update a department entity by ID",
                "operationId": "update-department",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Department ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Department entity",
                        "name": "department",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/ent.Department"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ent.Department"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            },
            "delete": {
                "description": "get department by ID",
                "produces": [
                    "application/json"
                ],
                "summary": "Delete a department entity by ID",
                "operationId": "delete-department",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Department ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/doctors": {
            "get": {
                "description": "list doctor entities",
                "produces": [
                    "application/json"
                ],
                "summary": "List doctor entities",
                "operationId": "list-doctor",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Offset",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/ent.Doctor"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            },
            "post": {
                "description": "Create doctor",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create doctor",
                "operationId": "create-doctor",
                "parameters": [
                    {
                        "description": "Doctor entity",
                        "name": "doctor",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.Doctor"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ent.Doctor"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/doctors/{id}": {
            "put": {
                "description": "update doctor by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Update a doctor entity by ID",
                "operationId": "update-doctor",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Doctor ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Doctor entity",
                        "name": "doctor",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/ent.Doctor"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ent.Doctor"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            },
            "delete": {
                "description": "get doctor by ID",
                "produces": [
                    "application/json"
                ],
                "summary": "Delete a doctor entity by ID",
                "operationId": "delete-doctor",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Doctor ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/nametitles": {
            "get": {
                "description": "list nametitle entities",
                "produces": [
                    "application/json"
                ],
                "summary": "List nametitle entities",
                "operationId": "list-nametitle",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Offset",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/ent.Nametitle"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            },
            "post": {
                "description": "Create nametitle",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create nametitle",
                "operationId": "create-nametitle",
                "parameters": [
                    {
                        "description": "Nametitle entity",
                        "name": "nametitle",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/ent.Nametitle"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ent.Nametitle"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/nametitles/{id}": {
            "get": {
                "description": "get nametitle by ID",
                "produces": [
                    "application/json"
                ],
                "summary": "Get a nametitle entity by ID",
                "operationId": "get-nametitle",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "NametitleID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ent.Nametitle"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            },
            "put": {
                "description": "update nametitle by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Update a nametitle entity by ID",
                "operationId": "update-nametitle",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Nametitle ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Nametitle entity",
                        "name": "nametitle",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/ent.Nametitle"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ent.Nametitle"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            },
            "delete": {
                "description": "get nametitle by ID",
                "produces": [
                    "application/json"
                ],
                "summary": "Delete a nametitle entity by ID",
                "operationId": "delete-nametitle",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Nametitle ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.Doctor": {
            "type": "object",
            "properties": {
                "degree": {
                    "type": "integer"
                },
                "department": {
                    "type": "integer"
                },
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "nametitle": {
                    "type": "integer"
                },
                "password": {
                    "type": "string"
                },
                "tel": {
                    "type": "string"
                }
            }
        },
        "ent.Degree": {
            "type": "object",
            "properties": {
                "Degree": {
                    "description": "Degree holds the value of the \"Degree\" field.",
                    "type": "string"
                },
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the DegreeQuery when eager-loading is set.",
                    "type": "object",
                    "$ref": "#/definitions/ent.DegreeEdges"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "integer"
                }
            }
        },
        "ent.DegreeEdges": {
            "type": "object",
            "properties": {
                "doctor": {
                    "description": "Doctor holds the value of the doctor edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Doctor"
                    }
                }
            }
        },
        "ent.Department": {
            "type": "object",
            "properties": {
                "Department": {
                    "description": "Department holds the value of the \"Department\" field.",
                    "type": "string"
                },
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the DepartmentQuery when eager-loading is set.",
                    "type": "object",
                    "$ref": "#/definitions/ent.DepartmentEdges"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "integer"
                }
            }
        },
        "ent.DepartmentEdges": {
            "type": "object",
            "properties": {
                "doctor": {
                    "description": "Doctor holds the value of the doctor edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Doctor"
                    }
                }
            }
        },
        "ent.Doctor": {
            "type": "object",
            "properties": {
                "Email": {
                    "description": "Email holds the value of the \"Email\" field.",
                    "type": "string"
                },
                "Name": {
                    "description": "Name holds the value of the \"Name\" field.",
                    "type": "string"
                },
                "Password": {
                    "description": "Password holds the value of the \"Password\" field.",
                    "type": "string"
                },
                "Tel": {
                    "description": "Tel holds the value of the \"Tel\" field.",
                    "type": "string"
                },
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the DoctorQuery when eager-loading is set.",
                    "type": "object",
                    "$ref": "#/definitions/ent.DoctorEdges"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "integer"
                }
            }
        },
        "ent.DoctorEdges": {
            "type": "object",
            "properties": {
                "degree": {
                    "description": "Degree holds the value of the degree edge.",
                    "type": "object",
                    "$ref": "#/definitions/ent.Degree"
                },
                "department": {
                    "description": "Department holds the value of the department edge.",
                    "type": "object",
                    "$ref": "#/definitions/ent.Department"
                },
                "nametitle": {
                    "description": "Nametitle holds the value of the nametitle edge.",
                    "type": "object",
                    "$ref": "#/definitions/ent.Nametitle"
                }
            }
        },
        "ent.Nametitle": {
            "type": "object",
            "properties": {
                "NameTitle": {
                    "description": "NameTitle holds the value of the \"NameTitle\" field.",
                    "type": "string"
                },
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the NametitleQuery when eager-loading is set.",
                    "type": "object",
                    "$ref": "#/definitions/ent.NametitleEdges"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "integer"
                }
            }
        },
        "ent.NametitleEdges": {
            "type": "object",
            "properties": {
                "doctor": {
                    "description": "Doctor holds the value of the doctor edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Doctor"
                    }
                }
            }
        },
        "gin.H": {
            "type": "object",
            "additionalProperties": true
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        },
        "BasicAuth": {
            "type": "basic"
        },
        "OAuth2AccessCode": {
            "type": "oauth2",
            "flow": "accessCode",
            "authorizationUrl": "https://example.com/oauth/authorize",
            "tokenUrl": "https://example.com/oauth/token",
            "scopes": {
                "admin": " Grants read and write access to administrative information"
            }
        },
        "OAuth2Application": {
            "type": "oauth2",
            "flow": "application",
            "tokenUrl": "https://example.com/oauth/token",
            "scopes": {
                "admin": " Grants read and write access to administrative information",
                "write": " Grants write access"
            }
        },
        "OAuth2Implicit": {
            "type": "oauth2",
            "flow": "implicit",
            "authorizationUrl": "https://example.com/oauth/authorize",
            "scopes": {
                "admin": " Grants read and write access to administrative information",
                "write": " Grants write access"
            }
        },
        "OAuth2Password": {
            "type": "oauth2",
            "flow": "password",
            "tokenUrl": "https://example.com/oauth/token",
            "scopes": {
                "admin": " Grants read and write access to administrative information",
                "read": " Grants read access",
                "write": " Grants write access"
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
	Version:     "1.0",
	Host:        "localhost:8080",
	BasePath:    "/api/v1",
	Schemes:     []string{},
	Title:       "SUT SA Example API Doctor",
	Description: "This is a sample server for SUT SE 2563",
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
