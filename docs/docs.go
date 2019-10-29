// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-10-29 11:12:06.056042 +0800 CST m=+0.062372008

package docs

import (
	"bytes"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "CatDogs API 文档",
        "title": "CatDogs API",
        "contact": {
            "name": "Yoko"
        },
        "license": {},
        "version": "1.0"
    },
    "host": "{{.Host}}",
    "basePath": "/api",
    "paths": {
        "/login": {
            "post": {
                "description": "登录接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "登录接口",
                "parameters": [
                    {
                        "description": "Login Params",
                        "name": "params",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/user.User"
                        }
                    }
                ],
                "responses": {
                    "0": {
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/libs.R"
                        }
                    },
                    "1002": {
                        "description": "用户不存在",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "1003": {
                        "description": "密码错误",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "999": {
                        "description": "服务器出问题",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "注册接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "注册接口",
                "parameters": [
                    {
                        "description": "Register Params",
                        "name": "params",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/user.User"
                        }
                    }
                ],
                "responses": {
                    "0": {
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/libs.R"
                        }
                    },
                    "1000": {
                        "description": "用户已存在",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "999": {
                        "description": "服务器出问题",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/setpost": {
            "post": {
                "description": "发布文章",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "文章"
                ],
                "summary": "发布文章",
                "parameters": [
                    {
                        "description": "SetPost Params",
                        "name": "params",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/post.SetPost"
                        }
                    }
                ],
                "responses": {
                    "0": {
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/libs.R"
                        }
                    },
                    "3000": {
                        "description": "参数错误",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "999": {
                        "description": "服务器出问题",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "libs.R": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "string"
                },
                "msg": {
                    "type": "string"
                }
            }
        },
        "post.SetPost": {
            "type": "object",
            "required": [
                "author",
                "content",
                "title"
            ],
            "properties": {
                "author": {
                    "type": "string"
                },
                "content": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "user.User": {
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
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo swaggerInfo

type s struct{}

func (s *s) ReadDoc() string {
	t, err := template.New("swagger_info").Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, SwaggerInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
