// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support",
            "url": "https://github.com/mhrndiva",
            "email": "714220050@std.ulbi.ac.id"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/insert": {
            "post": {
                "description": "Input data presensi.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Presensi"
                ],
                "summary": "Insert data presensi.",
                "parameters": [
                    {
                        "description": "Payload Body [RAW]",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.ReqPresensi"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Presensi"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/presensi": {
            "get": {
                "description": "Mengambil semua data presensi.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Presensi"
                ],
                "summary": "Get All Data Presensi.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Presensi"
                        }
                    }
                }
            }
        },
        "/presensi/{id}": {
            "get": {
                "description": "Ambil per ID data presensi.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Presensi"
                ],
                "summary": "Get By ID Data Presensi.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Masukan ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Presensi"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/update/{id}": {
            "put": {
                "description": "Ubah data presensi.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Presensi"
                ],
                "summary": "Update data presensi.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Masukan ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Payload Body [RAW]",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.ReqPresensi"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Presensi"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.Mahasiswa": {
            "type": "object",
            "properties": {
                "_id": {
                    "type": "string"
                },
                "alamat": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "jurusan": {
                    "type": "string"
                },
                "nama": {
                    "type": "string"
                },
                "npm": {
                    "type": "integer"
                },
                "phone_number": {
                    "type": "string"
                }
            }
        },
        "controller.Matkul": {
            "type": "object",
            "properties": {
                "dosen": {
                    "type": "string"
                },
                "jadwal": {
                    "type": "string"
                },
                "namamatkul": {
                    "type": "string"
                },
                "sks": {
                    "type": "integer"
                }
            }
        },
        "controller.Presensi": {
            "type": "object",
            "properties": {
                "_id": {
                    "type": "string"
                },
                "biodata": {
                    "$ref": "#/definitions/controller.Mahasiswa"
                },
                "checkin": {
                    "type": "string"
                },
                "datetime": {
                    "type": "string"
                },
                "matkul": {
                    "$ref": "#/definitions/controller.Matkul"
                },
                "npm": {
                    "type": "integer",
                    "example": 714220050
                }
            }
        },
        "controller.ReqMahasiswa": {
            "type": "object",
            "properties": {
                "alamat": {
                    "type": "string",
                    "example": "Sarijadi"
                },
                "email": {
                    "type": "string",
                    "example": "devi443@gmail.com"
                },
                "jurusan": {
                    "type": "string",
                    "example": "Informatika"
                },
                "nama": {
                    "type": "string",
                    "example": "Devi Wulandari"
                },
                "npm": {
                    "type": "integer",
                    "example": 714220050
                },
                "phone_number": {
                    "type": "string",
                    "example": "08123456789"
                }
            }
        },
        "controller.ReqMatkul": {
            "type": "object",
            "properties": {
                "dosen": {
                    "type": "string",
                    "example": "Roni Habibie"
                },
                "jadwal": {
                    "type": "string",
                    "example": "Senin, Selasa, Rabu, Kamis, Jumat, Sabtu, Minggu"
                },
                "namamatkul": {
                    "type": "string",
                    "example": "Kewirausahaan"
                },
                "sks": {
                    "type": "integer",
                    "example": 2
                }
            }
        },
        "controller.ReqPresensi": {
            "type": "object",
            "properties": {
                "biodata": {
                    "$ref": "#/definitions/controller.ReqMahasiswa"
                },
                "checkin": {
                    "type": "string",
                    "example": "Hadir"
                },
                "matkul": {
                    "$ref": "#/definitions/controller.ReqMatkul"
                },
                "npm": {
                    "type": "integer",
                    "example": 714220050
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "ws-ats-714220050-cc57ecdf5b73.herokuapp.com",
	BasePath:         "/",
	Schemes:          []string{"https", "http"},
	Title:            "TES SWAGGER ULBI",
	Description:      "This is a sample swagger for Fiber",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
