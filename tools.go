//go:build tools

//go:generate go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen -package main -o client.gen.go swagger.json

package main

import (
	_ "github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen"
)
