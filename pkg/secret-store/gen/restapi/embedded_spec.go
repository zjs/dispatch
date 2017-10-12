///////////////////////////////////////////////////////////////////////
// Copyright (C) 2016 VMware, Inc. All rights reserved.
// -- VMware Confidential
///////////////////////////////////////////////////////////////////////
// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

// SwaggerJSON embedded version of the swagger document used at generation time
var SwaggerJSON json.RawMessage

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/com.vmware.vs.secrets.v1+json"
  ],
  "produces": [
    "application/com.vmware.vs.secrets.v1+json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "An API for managing secrets for serverless functions.",
    "title": "Secret Store",
    "version": "0.0.1"
  },
  "basePath": "/v1/secret",
  "paths": {
    "/": {
      "get": {
        "tags": [
          "secret"
        ],
        "responses": {
          "200": {
            "description": "An array of registered secrets",
            "schema": {
              "$ref": "#/definitions/getOKBody"
            }
          },
          "default": {
            "description": "Standard error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "post": {
        "consumes": [
          "application/json"
        ],
        "tags": [
          "secret"
        ],
        "parameters": [
          {
            "name": "secret",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/Secret"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "The created secret.",
            "schema": {
              "$ref": "#/definitions/Secret"
            }
          },
          "default": {
            "description": "Standard error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/{secretName}": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "secret"
        ],
        "responses": {
          "200": {
            "description": "The secret identified by the secretName",
            "schema": {
              "$ref": "#/definitions/Secret"
            }
          },
          "404": {
            "description": "Resource Not Found if no secret exists with the given name",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "put": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "secret"
        ],
        "parameters": [
          {
            "name": "secret",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/Secret"
            }
          },
          {
            "pattern": "^[\\w\\d\\-]+$",
            "type": "string",
            "name": "secretName",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "201": {
            "description": "The updated secret",
            "schema": {
              "$ref": "#/definitions/Secret"
            }
          },
          "404": {
            "description": "Resource Not Found if no secret exists with the given name",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "default": {
            "description": "generic error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "delete": {
        "tags": [
          "secret"
        ],
        "parameters": [
          {
            "pattern": "^[\\w\\d\\-]+$",
            "type": "string",
            "name": "secretName",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "Successful deletion"
          },
          "404": {
            "description": "Resource Not Found if no secret exists with the given name",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "default": {
            "description": "generic error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "parameters": [
        {
          "pattern": "^[\\w\\d\\-]+$",
          "type": "string",
          "description": "name of the secret to operate on",
          "name": "secretName",
          "in": "path",
          "required": true
        }
      ]
    }
  },
  "definitions": {
    "Error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "Secret": {
      "type": "object",
      "required": [
        "name"
      ],
      "properties": {
        "id": {
          "type": "string",
          "format": "uuid",
          "readOnly": true
        },
        "name": {
          "type": "string",
          "pattern": "^[\\w\\d\\-]+$"
        },
        "secrets": {
          "$ref": "#/definitions/SecretValue"
        }
      }
    },
    "SecretValue": {
      "type": "object",
      "additionalProperties": {
        "type": "string"
      }
    },
    "getOKBody": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/Secret"
      },
      "x-go-gen-location": "operations"
    }
  },
  "tags": [
    {
      "description": "Operations on secrets",
      "name": "secret"
    }
  ]
}`))
}
