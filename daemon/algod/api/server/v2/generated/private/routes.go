// Package private provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/algorand/oapi-codegen DO NOT EDIT.
package private

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"github.com/algorand/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (POST /v2/register-participation-keys/{address})
	RegisterParticipationKeys(ctx echo.Context, address string, params RegisterParticipationKeysParams) error

	// (POST /v2/shutdown)
	ShutdownNode(ctx echo.Context, params ShutdownNodeParams) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// RegisterParticipationKeys converts echo context to params.
func (w *ServerInterfaceWrapper) RegisterParticipationKeys(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"fee":              true,
		"key-dilution":     true,
		"round-last-valid": true,
		"no-wait":          true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "address" -------------
	var address string

	err = runtime.BindStyledParameter("simple", false, "address", ctx.Param("address"), &address)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter address: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params RegisterParticipationKeysParams
	// ------------- Optional query parameter "fee" -------------
	if paramValue := ctx.QueryParam("fee"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "fee", ctx.QueryParams(), &params.Fee)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter fee: %s", err))
	}

	// ------------- Optional query parameter "key-dilution" -------------
	if paramValue := ctx.QueryParam("key-dilution"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "key-dilution", ctx.QueryParams(), &params.KeyDilution)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter key-dilution: %s", err))
	}

	// ------------- Optional query parameter "round-last-valid" -------------
	if paramValue := ctx.QueryParam("round-last-valid"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "round-last-valid", ctx.QueryParams(), &params.RoundLastValid)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter round-last-valid: %s", err))
	}

	// ------------- Optional query parameter "no-wait" -------------
	if paramValue := ctx.QueryParam("no-wait"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "no-wait", ctx.QueryParams(), &params.NoWait)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter no-wait: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.RegisterParticipationKeys(ctx, address, params)
	return err
}

// ShutdownNode converts echo context to params.
func (w *ServerInterfaceWrapper) ShutdownNode(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"timeout": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params ShutdownNodeParams
	// ------------- Optional query parameter "timeout" -------------
	if paramValue := ctx.QueryParam("timeout"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "timeout", ctx.QueryParams(), &params.Timeout)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter timeout: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ShutdownNode(ctx, params)
	return err
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}, si ServerInterface, m ...echo.MiddlewareFunc) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.POST("/v2/register-participation-keys/:address", wrapper.RegisterParticipationKeys, m...)
	router.POST("/v2/shutdown", wrapper.ShutdownNode, m...)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
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

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}
