// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	. "github.com/Peyton232/HTTPServerGenerator-exampleServer/models"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Returns all pets
	// (GET /pets)
	FindPets(ctx echo.Context, params FindPetsParams) error
	// Creates a new pet
	// (POST /pets)
	AddPet(ctx echo.Context) error
	// Update an existing pet
	// (PUT /pets)
	UpdatePet(ctx echo.Context) error
	// Deletes a pet by ID
	// (DELETE /pets/{id})
	DeletePet(ctx echo.Context, id int64) error
	// Returns a pet by ID
	// (GET /pets/{id})
	FindPetByID(ctx echo.Context, id int64) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// FindPets converts echo context to params.
func (w *ServerInterfaceWrapper) FindPets(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params FindPetsParams
	// ------------- Optional query parameter "tags" -------------

	err = runtime.BindQueryParameter("form", true, false, "tags", ctx.QueryParams(), &params.Tags)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter tags: %s", err))
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.FindPets(ctx, params)
	return err
}

// AddPet converts echo context to params.
func (w *ServerInterfaceWrapper) AddPet(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AddPet(ctx)
	return err
}

// UpdatePet converts echo context to params.
func (w *ServerInterfaceWrapper) UpdatePet(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.UpdatePet(ctx)
	return err
}

// DeletePet converts echo context to params.
func (w *ServerInterfaceWrapper) DeletePet(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeletePet(ctx, id)
	return err
}

// FindPetByID converts echo context to params.
func (w *ServerInterfaceWrapper) FindPetByID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.FindPetByID(ctx, id)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/pets", wrapper.FindPets)
	router.POST(baseURL+"/pets", wrapper.AddPet)
	router.PUT(baseURL+"/pets", wrapper.UpdatePet)
	router.DELETE(baseURL+"/pets/:id", wrapper.DeletePet)
	router.GET(baseURL+"/pets/:id", wrapper.FindPetByID)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"/3l4zuLaEwAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
