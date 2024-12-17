// Code generated by swaggo/swag. DO NOT EDIT.

package swagger

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support",
            "email": "andreysapozhkov535@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/segment": {
            "post": {
                "description": "add new segment",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Segment API"
                ],
                "summary": "Add new segment.",
                "parameters": [
                    {
                        "description": "Segment data",
                        "name": "SegmentDTO",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_Inspirate789_backend-trainee-assignment-2023_internal_segment_usecase_dto.SegmentDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "delete segment",
                "tags": [
                    "Segment API"
                ],
                "summary": "Delete segment.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Segment name",
                        "name": "name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user": {
            "post": {
                "description": "add new user",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "User API"
                ],
                "summary": "Add new user.",
                "parameters": [
                    {
                        "description": "User data",
                        "name": "UserDTO",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_Inspirate789_backend-trainee-assignment-2023_internal_user_usecase_dto.UserDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "delete user",
                "tags": [
                    "User API"
                ],
                "summary": "Delete user.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/history": {
            "get": {
                "description": "get the history of changing user segments; returns the web link to csv file with report",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User API"
                ],
                "summary": "Get the history of changing user segments.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Year and month in history",
                        "name": "year_month",
                        "in": "query",
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
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/segments": {
            "get": {
                "description": "get user segments",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User API"
                ],
                "summary": "Get user segments.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_Inspirate789_backend-trainee-assignment-2023_internal_user_usecase_dto.UserSegmentsOutputDTO"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "patch": {
                "description": "change user segments",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "User API"
                ],
                "summary": "Change user segments.",
                "parameters": [
                    {
                        "description": "Old and new user segments",
                        "name": "UserSegmentsInputDTO",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_Inspirate789_backend-trainee-assignment-2023_internal_user_usecase_dto.UserSegmentsInputDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "github_com_Inspirate789_backend-trainee-assignment-2023_internal_segment_usecase_dto.SegmentDTO": {
            "type": "object",
            "properties": {
                "name": {
                    "description": "Name\nrequired: true\nmin length: 1\nexample: \"AVITO_VOICE_MESSAGES\"",
                    "type": "string"
                },
                "ttl": {
                    "description": "TTL - segment existing time (in hours)\nrequired: false\nmin: 1\nexample: 72",
                    "type": "integer"
                },
                "user_percentage": {
                    "description": "UserPercentage - part of all users that segment contains (in %)\nrequired: false\nmin: 0\nmax: 100\nexample: 50",
                    "type": "number"
                }
            }
        },
        "github_com_Inspirate789_backend-trainee-assignment-2023_internal_user_usecase_dto.UserDTO": {
            "type": "object",
            "properties": {
                "user_id": {
                    "description": "UserID\nrequired: true\nmin: 1\nexample: 75",
                    "type": "integer"
                }
            }
        },
        "github_com_Inspirate789_backend-trainee-assignment-2023_internal_user_usecase_dto.UserSegmentsInputDTO": {
            "type": "object",
            "properties": {
                "new_segment_names": {
                    "description": "NewSegmentNames - segment names to adding\nrequired: false\nmin items: 0\nexample: [\"AVITO_VOICE_MESSAGES\", \"AVITO_DISCOUNT_50\"]",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "old_segment_names": {
                    "description": "OldSegmentNames - segment names to removing\nrequired: false\nmin items: 0\nexample: [\"AVITO_VOICE_MESSAGES\", \"AVITO_PERFORMANCE_VAS\", \"AVITO_DISCOUNT_30\"]",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "ttl": {
                    "description": "TTL - segment existing time (in hours)\nrequired: false\nmin: 1\nexample: 72",
                    "type": "integer"
                },
                "user_id": {
                    "description": "UserID\nrequired: true\nmin: 1\nexample: 75",
                    "type": "integer"
                }
            }
        },
        "github_com_Inspirate789_backend-trainee-assignment-2023_internal_user_usecase_dto.UserSegmentsOutputDTO": {
            "type": "object",
            "properties": {
                "segment_names": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{"http"},
	Title:            "Application API",
	Description:      "This is an application API.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}