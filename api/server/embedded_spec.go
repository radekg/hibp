// Code generated by go-swagger; DO NOT EDIT.

package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "text/plain"
  ],
  "schemes": [
    "http",
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Self-hosted HIBP password hash checker",
    "title": "Self-hosted HIBP password hash checker",
    "version": "latest"
  },
  "basePath": "/",
  "paths": {
    "/range/{hashPrefix}": {
      "get": {
        "description": "Search password hashes by range.",
        "produces": [
          "text/plain"
        ],
        "schemes": [
          "http",
          "https"
        ],
        "tags": [
          "range"
        ],
        "operationId": "rangeSearch",
        "parameters": [
          {
            "type": "string",
            "name": "hashPrefix",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Request was processed successfully.",
            "schema": {
              "type": "string"
            }
          },
          "400": {
            "description": "Hash prefix must be exactly 5 characters long.",
            "schema": {
              "type": "string"
            }
          },
          "404": {
            "description": "No results found."
          },
          "500": {
            "description": "Server encountered an error.",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "text/plain"
  ],
  "schemes": [
    "http",
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Self-hosted HIBP password hash checker",
    "title": "Self-hosted HIBP password hash checker",
    "version": "latest"
  },
  "basePath": "/",
  "paths": {
    "/range/{hashPrefix}": {
      "get": {
        "description": "Search password hashes by range.",
        "produces": [
          "text/plain"
        ],
        "schemes": [
          "http",
          "https"
        ],
        "tags": [
          "range"
        ],
        "operationId": "rangeSearch",
        "parameters": [
          {
            "type": "string",
            "name": "hashPrefix",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Request was processed successfully.",
            "schema": {
              "type": "string"
            }
          },
          "400": {
            "description": "Hash prefix must be exactly 5 characters long.",
            "schema": {
              "type": "string"
            }
          },
          "404": {
            "description": "No results found."
          },
          "500": {
            "description": "Server encountered an error.",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    }
  }
}`))
}
