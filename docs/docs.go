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
            "name": "Archivista Contributors",
            "url": "https://github.com/in-toto/archivista/issues/new"
        },
        "license": {
            "url": "https://opensource.org/licenses/Apache-2"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/donwload/{gitoid}": {
            "post": {
                "description": "download an attestation",
                "produces": [
                    "application/json"
                ],
                "summary": "Download",
                "deprecated": true,
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dsse.Envelope"
                        }
                    }
                }
            }
        },
        "/upload": {
            "post": {
                "description": "stores an attestation",
                "produces": [
                    "application/json"
                ],
                "summary": "Store",
                "deprecated": true,
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.StoreResponse"
                        }
                    }
                }
            }
        },
        "/v1/donwload/{gitoid}": {
            "post": {
                "description": "download an attestation",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "attestation"
                ],
                "summary": "Download",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dsse.Envelope"
                        }
                    }
                }
            }
        },
        "/v1/query": {
            "post": {
                "description": "GraphQL query",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "graphql"
                ],
                "summary": "Query GraphQL",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/archivista.Resolver"
                        }
                    }
                }
            }
        },
        "/v1/upload": {
            "post": {
                "description": "stores an attestation",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "attestation"
                ],
                "summary": "Store",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.StoreResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.StoreResponse": {
            "type": "object",
            "properties": {
                "gitoid": {
                    "type": "string"
                }
            }
        },
        "archivista.Resolver": {
            "type": "object"
        },
        "dsse.Envelope": {
            "type": "object",
            "properties": {
                "payload": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "payloadType": {
                    "type": "string"
                },
                "signatures": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dsse.Signature"
                    }
                }
            }
        },
        "dsse.Signature": {
            "type": "object",
            "properties": {
                "certificate": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "intermediates": {
                    "type": "array",
                    "items": {
                        "type": "array",
                        "items": {
                            "type": "integer"
                        }
                    }
                },
                "keyid": {
                    "type": "string"
                },
                "sig": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "timestamps": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dsse.SignatureTimestamp"
                    }
                }
            }
        },
        "dsse.SignatureTimestamp": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "type": {
                    "$ref": "#/definitions/dsse.SignatureTimestampType"
                }
            }
        },
        "dsse.SignatureTimestampType": {
            "type": "string",
            "enum": [
                "tsp"
            ],
            "x-enum-varnames": [
                "TimestampRFC3161"
            ]
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "v1",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Archivista API",
	Description:      "Archivista API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
