// Package generated provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/algorand/oapi-codegen DO NOT EDIT.
package generated

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
	// Get account information.
	// (GET /v2/accounts/{address})
	AccountInformation(ctx echo.Context, address string, params AccountInformationParams) error
	// Get a list of unconfirmed transactions currently in the transaction pool by address.
	// (GET /v2/accounts/{address}/transactions/pending)
	GetPendingTransactionsByAddress(ctx echo.Context, address string, params GetPendingTransactionsByAddressParams) error
	// Get application information.
	// (GET /v2/applications/{application-id})
	GetApplicationByID(ctx echo.Context, applicationId uint64) error
	// Get asset information.
	// (GET /v2/assets/{asset-id})
	GetAssetByID(ctx echo.Context, assetId uint64) error
	// Get the block for the given round.
	// (GET /v2/blocks/{round})
	GetBlock(ctx echo.Context, round uint64, params GetBlockParams) error
	// Get the current supply reported by the ledger.
	// (GET /v2/ledger/supply)
	GetSupply(ctx echo.Context) error
	// Gets the current node status.
	// (GET /v2/status)
	GetStatus(ctx echo.Context) error
	// Gets the node status after waiting for the given round.
	// (GET /v2/status/wait-for-block-after/{round})
	WaitForBlock(ctx echo.Context, round uint64) error
	// Compile TEAL source code to binary, produce its hash
	// (POST /v2/teal/compile)
	TealCompile(ctx echo.Context) error
	// Provide debugging information for a transaction (or group).
	// (POST /v2/teal/dryrun)
	TealDryrun(ctx echo.Context) error
	// Broadcasts a raw transaction to the network.
	// (POST /v2/transactions)
	RawTransaction(ctx echo.Context) error
	// Get parameters for constructing a new transaction
	// (GET /v2/transactions/params)
	TransactionParams(ctx echo.Context) error
	// Get a list of unconfirmed transactions currently in the transaction pool.
	// (GET /v2/transactions/pending)
	GetPendingTransactions(ctx echo.Context, params GetPendingTransactionsParams) error
	// Get a specific pending transaction.
	// (GET /v2/transactions/pending/{txid})
	PendingTransactionInformation(ctx echo.Context, txid string, params PendingTransactionInformationParams) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// AccountInformation converts echo context to params.
func (w *ServerInterfaceWrapper) AccountInformation(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
		"format": true,
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
	var params AccountInformationParams
	// ------------- Optional query parameter "format" -------------
	if paramValue := ctx.QueryParam("format"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "format", ctx.QueryParams(), &params.Format)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter format: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AccountInformation(ctx, address, params)
	return err
}

// GetPendingTransactionsByAddress converts echo context to params.
func (w *ServerInterfaceWrapper) GetPendingTransactionsByAddress(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
		"max":    true,
		"format": true,
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
	var params GetPendingTransactionsByAddressParams
	// ------------- Optional query parameter "max" -------------
	if paramValue := ctx.QueryParam("max"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "max", ctx.QueryParams(), &params.Max)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter max: %s", err))
	}

	// ------------- Optional query parameter "format" -------------
	if paramValue := ctx.QueryParam("format"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "format", ctx.QueryParams(), &params.Format)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter format: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetPendingTransactionsByAddress(ctx, address, params)
	return err
}

// GetApplicationByID converts echo context to params.
func (w *ServerInterfaceWrapper) GetApplicationByID(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "application-id" -------------
	var applicationId uint64

	err = runtime.BindStyledParameter("simple", false, "application-id", ctx.Param("application-id"), &applicationId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter application-id: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetApplicationByID(ctx, applicationId)
	return err
}

// GetAssetByID converts echo context to params.
func (w *ServerInterfaceWrapper) GetAssetByID(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "asset-id" -------------
	var assetId uint64

	err = runtime.BindStyledParameter("simple", false, "asset-id", ctx.Param("asset-id"), &assetId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter asset-id: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetAssetByID(ctx, assetId)
	return err
}

// GetBlock converts echo context to params.
func (w *ServerInterfaceWrapper) GetBlock(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
		"format": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "round" -------------
	var round uint64

	err = runtime.BindStyledParameter("simple", false, "round", ctx.Param("round"), &round)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter round: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetBlockParams
	// ------------- Optional query parameter "format" -------------
	if paramValue := ctx.QueryParam("format"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "format", ctx.QueryParams(), &params.Format)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter format: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetBlock(ctx, round, params)
	return err
}

// GetSupply converts echo context to params.
func (w *ServerInterfaceWrapper) GetSupply(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetSupply(ctx)
	return err
}

// GetStatus converts echo context to params.
func (w *ServerInterfaceWrapper) GetStatus(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetStatus(ctx)
	return err
}

// WaitForBlock converts echo context to params.
func (w *ServerInterfaceWrapper) WaitForBlock(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "round" -------------
	var round uint64

	err = runtime.BindStyledParameter("simple", false, "round", ctx.Param("round"), &round)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter round: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.WaitForBlock(ctx, round)
	return err
}

// TealCompile converts echo context to params.
func (w *ServerInterfaceWrapper) TealCompile(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.TealCompile(ctx)
	return err
}

// TealDryrun converts echo context to params.
func (w *ServerInterfaceWrapper) TealDryrun(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.TealDryrun(ctx)
	return err
}

// RawTransaction converts echo context to params.
func (w *ServerInterfaceWrapper) RawTransaction(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.RawTransaction(ctx)
	return err
}

// TransactionParams converts echo context to params.
func (w *ServerInterfaceWrapper) TransactionParams(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.TransactionParams(ctx)
	return err
}

// GetPendingTransactions converts echo context to params.
func (w *ServerInterfaceWrapper) GetPendingTransactions(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
		"max":    true,
		"format": true,
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
	var params GetPendingTransactionsParams
	// ------------- Optional query parameter "max" -------------
	if paramValue := ctx.QueryParam("max"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "max", ctx.QueryParams(), &params.Max)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter max: %s", err))
	}

	// ------------- Optional query parameter "format" -------------
	if paramValue := ctx.QueryParam("format"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "format", ctx.QueryParams(), &params.Format)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter format: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetPendingTransactions(ctx, params)
	return err
}

// PendingTransactionInformation converts echo context to params.
func (w *ServerInterfaceWrapper) PendingTransactionInformation(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
		"format": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "txid" -------------
	var txid string

	err = runtime.BindStyledParameter("simple", false, "txid", ctx.Param("txid"), &txid)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter txid: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params PendingTransactionInformationParams
	// ------------- Optional query parameter "format" -------------
	if paramValue := ctx.QueryParam("format"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "format", ctx.QueryParams(), &params.Format)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter format: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PendingTransactionInformation(ctx, txid, params)
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

	router.GET("/v2/accounts/:address", wrapper.AccountInformation, m...)
	router.GET("/v2/accounts/:address/transactions/pending", wrapper.GetPendingTransactionsByAddress, m...)
	router.GET("/v2/applications/:application-id", wrapper.GetApplicationByID, m...)
	router.GET("/v2/assets/:asset-id", wrapper.GetAssetByID, m...)
	router.GET("/v2/blocks/:round", wrapper.GetBlock, m...)
	router.GET("/v2/ledger/supply", wrapper.GetSupply, m...)
	router.GET("/v2/status", wrapper.GetStatus, m...)
	router.GET("/v2/status/wait-for-block-after/:round", wrapper.WaitForBlock, m...)
	router.POST("/v2/teal/compile", wrapper.TealCompile, m...)
	router.POST("/v2/teal/dryrun", wrapper.TealDryrun, m...)
	router.POST("/v2/transactions", wrapper.RawTransaction, m...)
	router.GET("/v2/transactions/params", wrapper.TransactionParams, m...)
	router.GET("/v2/transactions/pending", wrapper.GetPendingTransactions, m...)
	router.GET("/v2/transactions/pending/:txid", wrapper.PendingTransactionInformation, m...)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+x9/XPbOJLov4LTXdUkOVFyvmY3rpq650nmw2+TTCr27O27OG8XIlsS1iTAAUBLmjz/",
	"76/QAEiQBCX5I1+z/imxCDQaje5Gd6PR+DBKRVEKDlyr0eGHUUklLUCDxL9omoqK64Rl5q8MVCpZqZng",
	"o0P/jSgtGV+MxiNmfi2pXo7GI04LaNqY/uORhN8qJiEbHWpZwXik0iUU1ADWm9K0riGtk4VIHIgjC+L4",
	"xehyyweaZRKU6mP5C883hPE0rzIgWlKuaGo+KbJiekn0kiniOhPGieBAxJzoZasxmTPIMzXxk/ytArkJ",
	"ZukGH57SZYNiIkUOfTyfi2LGOHisoEaqXhCiBclgjo2WVBMzgsHVN9SCKKAyXZK5kDtQtUiE+AKvitHh",
	"u5ECnoHE1UqBXeB/5xLgd0g0lQvQo/fj2OTmGmSiWRGZ2rGjvgRV5VoRbItzXLAL4MT0mpBXldJkBoRy",
	"8vbH5+Tx48fPzEQKqjVkjskGZ9WMHs7Jdh8djjKqwX/u8xrNF0JSniV1+7c/PsfxT9wE921FlYK4sByZ",
	"L+T4xdAEfMcICzGuYYHr0OJ+0yMiFM3PM5gLCXuuiW18q4sSjv9ZVyWlOl2WgnEdWReCX4n9HNVhQfdt",
	"OqxGoNW+NJSSBui7g+TZ+w8Pxw8PLv/93VHyP+7Pp48v95z+8xruDgpEG6aVlMDTTbKQQFFalpT36fHW",
	"8YNaiirPyJJe4OLTAlW960tMX6s6L2heGT5hqRRH+UIoQh0bZTCnVa6JH5hUPDdqykBz3E6YIqUUFyyD",
	"bGy072rJ0iVJqbIgsB1ZsTw3PFgpyIZ4LT67LcJ0GZLE4HUteuCEvlxiNPPaQQlYozZI0lwoSLTYsT35",
	"HYfyjIQbSrNXqattVuR0CQQHNx/sZou044an83xDNK5rRqgilPitaUzYnGxERVa4ODk7x/5uNoZqBTFE",
	"w8Vp7aNGeIfI1yNGhHgzIXKgHInn5a5PMj5ni0qCIqsl6KXb8ySoUnAFRMz+Cak2y/6/T355TYQkr0Ap",
	"uoA3ND0nwFORDa+xGzS2g/9TCbPghVqUND2Pb9c5K1gE5Vd0zYqqILwqZiDNevn9QQsiQVeSDyFkIe7g",
	"s4Ku+4OeyoqnuLjNsC1DzbASU2VONxNyPCcFXX93MHboKELznJTAM8YXRK/5oJFmxt6NXiJFxbM9bBht",
	"FizYNVUJKZszyEgNZQsmbphd+DB+NXwayypAxwMZRKceZQc6HNYRnjGia76Qki4gYJkJ+dVpLvyqxTnw",
	"WsGR2QY/lRIumKhU3WkARxx6u3nNhYaklDBnER47ceQw2sO2ceq1cAZOKrimjENmNC8iLTRYTTSIUzDg",
	"dmemv0XPqIJvnwxt4M3XPVd/LrqrvnXF91ptbJRYkYzsi+arE9i42dTqv4fzF46t2CKxP/cWki1OzVYy",
	"ZzluM/806+fJUClUAi1C+I1HsQWnupJweMYfmL9IQk405RmVmfmlsD+9qnLNTtjC/JTbn16KBUtP2GKA",
	"mDWuUW8KuxX2HwMvro71Ouo0vBTivCrDCaUtr3S2IccvhhbZwrwqYx7VrmzoVZyuvadx1R56XS/kAJKD",
	"tCupaXgOGwkGW5rO8Z/1HPmJzuXvMWIaznU7LEYDXJTgrfvN/GRkHawzQMsyZyk11Jzivnn4IcDkPyTM",
	"R4ejf582IZKp/aqmDq4dsb1s96Ao9ea+mf5RA//2MWh6xrAIPhPG7XJh07F1Em8fHwM1iglarh0cvs9F",
	"en4tHEopSpCa2fWdGTh90UHwZAk0A0kyqumk8bKs4TUgANjxZ+yHbhPIyJ73C/6H5sR8NmJJtbfnjC3L",
	"lLHqRBB5yowJaDcWO5JpgKapIIW1+oix1q6E5fNmcKuxaxX7zpHlfRdaZHV+sIYmwR5+EmbqjRt5NBPy",
	"evzSYQROGueYUAO1NofNzNsri02rMnH0iRjYtkEHUBOP7OvZkEJd8PvQKpDshjonmn4E6igD9Tao0wb0",
	"qagjipLlcAvyvaRq2Z+csZAePyInPx89ffjo74+efmu2+FKKhaQFmW00KHLPbUxE6U0O9/szxo2iynUc",
	"+rdPvAvWhruTcohwDXsfup2C0SSWYsQGHAx2L+RGVvwWSAhSChkxmpGltEhFnlyAVExE4h9vXAviWhi9",
	"ZQ33zu8WW7Kiipix0Z+reAZyEqO8cdTQJtBQqF0biwV9uuYNbRxAKiXd9FbAzjcyOzfuPmvSJr53DxQp",
	"QSZ6zUkGs2oR7mlkLkVBKMmwIyrQ1yKDE011pW5BOzTAGmTMQoQo0JmoNKGEi8wIumkc1xsDwVCMwmDw",
	"SIeqSC/tfjUDY16ntFosNTF2qYgtbdMxoaldlAT3FjXgO9ZOv21lh7OBtlwCzTZkBsCJmDkHzbmOOEmK",
	"cR3tj2yc1mrQqp2KFl6lFCkoBVnizqd2oubPunCR9RYyId6Ibz0IUYLMqbwmrlpomu/AE9v0sVWN9eGc",
	"2j7W+w2/bf26g4erSKXxUS0TGFPHCHcOGoZIuJMmVTlwnuF2u1NWGJEgnHKhIBU8U1FgOVU62SUKplFr",
	"SzbLGnBfjPsR8IDX/pIqbf1mxjM026wI4zjYB4cYRnhQSxvIf/UKug87NbqHq0rV2lpVZSmkhiw2Bw7r",
	"LWO9hnU9lpgHsOstQQtSKdgFeYhKAXxHLDsTSyCqXeCmDiz1J4cxcqNbN1FStpBoCLENkRPfKqBuGNMd",
	"QMTY+HVPZBymOpxTB5LHI6VFWRqdpJOK1/2GyHRiWx/pX5u2feaiutGVmQAzuvY4OcxXlrI2mr+kxl5C",
	"yKSg50bfo/VjHfw+zkYYE8V4Csk2zjdieWJahSKwQ0gHDFJ3XhiM1hGODv9GmW6QCXaswtCEr2gdv7Hh",
	"6tMmlHMLBsIL0JTlqjYC6ph4MwqGz7upDcZik5AC1/nG8PCcycKeQOHeofxv1sTI3Cj2rKURS54RCSsq",
	"M9+i77G4gy6ewTqub6mLE2SwJiyO6LwejWmS+jMhd4g2ie8beIxjkVOxAz78YPixYKkU1J7bGcLbPUvX",
	"R1MSCmqwwxMkt8cOj8n4IrHHhJHdyn73x4g+fBsuVRyuX55BQatXZLUEPJkw2rNDxHCRjdcECoYmUgqR",
	"J7X/0A1C9/RMd6Rzlp5DRgxDotXj1N83bZzMIOSeWVRVh+lXy403qMoSOGT3J4QccYJC5JzYzlbXGZx/",
	"o7eNv8ZRswpPDCknOMnJGY/7ifa88YZc5MFs5x2bgHPDoSyQ7QPpNR9gILrCcLkBF+XIraGpE+wZ6Lae",
	"Kg+YymKxj/r8CbNSaGuVWYbWbqO+VDUrGKamBM3GRlf408K+u8T0hJBTlBZjriq4AGn8carsJu/O9gtm",
	"vB5VpSlAdnjGkxYmqSjcwPea/1pBPKsODh4DObjf7aO0sVOcZW5loNv3O3Iwtp+QXOQ7cjY6G/UgSSjE",
	"BWTWOwn52vbaCfbfarhn/JeeKiIF3Vi/xssiUdV8zlJmiZ4Lo8kWomNucIFfQBr0wHgHijA9RuWNFEUz",
	"za5LI4Dx7fE2HOgIVGOgmc1DSrrxZ0Rt3lEE1jQ1s6SoZDZkZRil5rP+LqdFmYQAonG+LSO6CKw9CfXR",
	"kWvKXTdOMh5Zd247fqcdh65FjoBdJ7uNth4xohjsI/5HpBRm1ZnLBvEpAzlTuoek8ywx/F4zZGTTmZD/",
	"IyqSUpTfstJQG/VCoqWMHpQZAXdRP6azTRoKQQ4FWH8bvzx40J34gwduzZkic1j5FCrTsEuOBw+sEAil",
	"bywBHdZcH0dMBoxymt00kva6pGo52RnxRLh7BToD0Mcv/IAoTErhFnM5HhlfK9/cgsBbQESCs3BUK+qg",
	"7FcxD9O13PqpjdJQ9ENntuvfB2yvt95F6O20gueMQ1IIDptohjLj8Ao/RvdpZJGBziisQ327LlQL/w5a",
	"7XH2Wc2b0hdXO2CJN3Xy2C0sfhduJ2oaJqqhlQl5SShJc4YRKcGVllWqzzhFD7ljBnXYwvv9wzGT575J",
	"PEgTiaE4UGecKkPD2m+ORtPnEImI/QjgQyeqWixAdcwiMgc4464V46TiTONYaFUmdsFKkHjsMbEtjSUw",
	"pzmGeH4HKcis0m3Vi/k01rKxIVwzDBHzM041yYEqTV4xfrpGcN7v8TzDQa+EPK+pELdbF8BBMZXET4Z+",
	"sl9/pmrpp28aemXjOtsopYHfJN1sNLQSdv/vvf86fHeU/A9Nfj9Inv3n9P2HJ5f3H/R+fHT53Xf/r/3T",
	"48vv7v/Xf8RWyuMey/ZwmB+/cGbJ8Qvce5robQ/3TxZ9LBhPokxm3IWCcUwa7PAWuWd2UM9A95s4sFv1",
	"M67X3DDSBc1ZRvX12KGr4nqyaKWjwzWthegEk/xc38fcnYVISpqe44HraMH0sppNUlFMvTk2XYjaNJtm",
	"FArB8Vs2pSWbGvd2evFwx9Z4A31FIuoK86nsSVqQDxMxS90RR8tDMhDtfQCbUGY8hBcwZ5yZ74dnPKOa",
	"TmdUsVRNKwXye5pTnsJkIcghcSBfUE3Rse7Eg4au7GC2s8OmrGY5S8l5uL81/D4UXzk7e2eofnb2vnc8",
	"0d+N3FBRxrcDJCuml6LSiYupDTvnTQADIdvwzrZRx8TBtsvsYnYOflz/0bJUSS5SmidKUw3x6ZdlbqYf",
	"7JmKYCfMhiFKC+k1i1E3LlBg1ve1cAc0kq58knJlnOF/FLR8x7h+TxLn1B6V5UsD88Tg8Q8nwEbrbkpo",
	"OTB75jE1wFTMe8GZWzPlyilSCPXE9vI3dVScdOYT0g7bGFlrovfXJZQB9bPIzepem04BjCh1Kr1MjFBF",
	"Z6UMb6FABHfL6MJoGH+kYpxRw33ursMMSLqE9BwyjBtj5G3c6u5PMp2+9jLLlL2eYDOhMIcWnawZkKrM",
	"qNvRKN90kxkVaO0zON/COWxORZOCe5XsxcvxyAWHE8MzQxJSGnoEqlXM2/LiA8ydxXehcQzgliVZ5GLm",
	"xKpmi8OaL3yfYQmy+v4WpCfGFDUZtvB7SWWEEJb5B0hwjYkaeDdi/dj0Sio1S1lp579fyuabVh8DZJdW",
	"j+pxMe+q6542japv2ziZURXX3GC+mPUwMtRNGvAj2XgFtYc6eMXVMe4sh+B0QjnJphJNCD9te2dvCLU4",
	"l4DkzXbq0WhTJNy3l+5UiV00Z0l4mrjPDrfzcMNwkT8GZu2gLjPj5nBBB+Prg7nlx8HZbnBlqc4c94qt",
	"Kwzj+haBvT3sM8x9WrnPJR+Nr5QXPh65FJ7YcgiO23sGOSyoCydjcpBjFIfaNypYIIPHL/O5cfpJEjsm",
	"pkqJlNkztUaXuzHAWH8PCLHhCrI3hBgbB2hjHA4Bk9cilE2+uAqSHBgG7qiHjRG84G/YHcdqrnE7u3Kn",
	"/dfXHY0QjZtrFnYZ+zGV8SiqkoZM81YrYpvMoOcgxFjUqKZ+lKEfy1CQA27HSUuzJuex2JOxKgDZ8MR3",
	"C+x1co/NzSZ/PwjHSlgYj7bxAo20+rDGp/XEL4SGZM6k0gk6oNHpmUY/KjQGfzRN4+qnRSpi74GyLK59",
	"cNhz2CQZy6v4artx//LCDPu6dlxUNTuHDW4yQNMlmeG9ZbMLtYY3bbYMbVMltk74pZ3wS3pr892Pl0xT",
	"M7AUQnfG+Eq4qqNPtglThAFjzNFftUGSblEv6Pu8gFzHss6DayLoThqFaa9LDLrrPWHKPOxt5leAxbDm",
	"tZCicwkM3a2zsAkkNkckuPbbT4UdkAFalixbd5xnC3UgSQIN+CsY6tbi71EBV9cB20GBxlGOZoZJ8M6+",
	"XdJgz7QXuHk4t/5aGx7Ey+y7JnUKNP8LbP5q2uK4o8vx6Ga+eYcoDSo14L1pEzGd3lAmO35ewDDhrwH9",
	"BjgnYlv6FblyvCPODBbejvm+qdkvygcYObYuaiu0d0WWoGUpxQXNE3cbYkh0pLhwooPN/eWJT7/BpzlQ",
	"aSNkW3HGduUXgrPxtmO5VqdB1AUtYe+XWyMvWLj6+lkYqPHJYC070TC8YwyrJerNMxQDF7iZxw+fdoZh",
	"7ADJXoIRVSYhgBtH/YKoaXKrWqonHXH+a1Z4h0yHY225zF7Yeg2KCN5NSTAmInqwyC4F3ZhVtFHfvnDz",
	"qkgMgycqZ2k8LMFnysgIrwpM8t9oINh4wNg0ECs2EJvnFQtgmWZqj7OdDpLBGFFiYshoC+1mwhXaqjj7",
	"rQLCMuDafJIuRaklLEY2fJ5pfzuI57Q6wC6ttQZ/ExvCgBqyHhCJ7QZEGEGOJBB7h9JPtA59mx+CwN8V",
	"ToDCEXtbypbTG8cfjpvt2fSyHQkO62L1dZBhDFtDYXdRLh+WWFpEB8aIFtka1NhHw9oac5X319ONWkZ0",
	"Q4Vss+lorkQETMVXlNuaOaafpaHrrcDGBEyvlZB4z0VB9EyZqWQuxe8Q91TnZqEiWVOOlJjvhL0nkfsD",
	"XSVaR12aamieviEeg6w9ZAkFH0n7hG5AwpHLg9A4poH6ABbllq1tfZ/WYWtcOMIEiamF3wiHw7mXVJLT",
	"1YzGbrYbk8XgdNQcwrRCbVoQ39mvgqqznx3vBec5dVtmL4eUIJvUxlszUL4uls8gZQXN45HXDKnfvh6Y",
	"sQWzRZIqBUEVHgfIVpezXOQqGdljroY0x3NyMA7qfLnVyNgFU2yWA7Z4aFvMqMJdqw6n1l3M9IDrpcLm",
	"j/Zovqx4JiHTS2UJqwSpjUh0u+rY9gz0CoCTA2z38Bm5h1F9xS7gvqGis0VGhw+fYRKF/eMgttm5amjb",
	"9EqGiuW/nWKJ8zEea1gYZpNyUCfRi0q2hOWwCtsiTbbrPrKELZ3W2y1LBeV0AfHT2mIHTrYvriYGBTt0",
	"4Zmtv6a0FBvCdHx80NTop4FEKqP+LBouu70wAqQFUaIw/NSU2LGDenC2mJurcuHx8h/xCKX0txQ6Duen",
	"9bXsXh6bNR50vaYFtMk6JtTe58OLFu4eqFOIk4HyAiAv4oPIgQX2+6brS+5xwZPCyE52v0nRC/gverte",
	"aJpHh9Ved3XTYraD3tfUMlCSQcJWLcLSQCddm8SVjM+TVmaoX9++dBtDIWTsqnyjDd0mIUFLBhdRie2m",
	"mtWWSb1deMrHDBRbUOCoLCOGt/+Eq6+8U07LMnDMfWUBktI8cnATbMx9Cy4bqCa7bSlvK6gZGHA7wpu+",
	"5MJvFSgdu/eEH2z6Enq2hiyOKMAz3GMnxN4TMgvbuumBexsrqtzeGoBsAdJRtypzQbMxMXBOfzh6Seyo",
	"yl1qxPspWO5hYe+c1UwUCZIF1/T3S2zwVaTiyU5XLJFhuCsCycxZabw5qzQtylhuqGlx6htgAuoFZblP",
	"J0CVH9JmQl7Y3VZ5XW4Hae4Wkno4J9/5QuBdbqo1TZe4jbWUvuW8qG+8d5USn16tgnJ/deW0+u67vTyo",
	"hS9UYuuUjIkwtsaKKVuyFC6gnY5a52Y7M8qnp7anJyvOLZ/E94QtdweuQ3aPnD2o82GgKGYdwl9RHyhR",
	"yRSuypEn2Ct6E6lbAaZX549DdrrmdVktX4o6pVxwluI9oKBIao2yK3+6T5xyjytTXRfVC7iTz4hwRevO",
	"1KkAjoqDlWi8GnSEG9gr7FezqJY77J8a62wa52sBWjm9BtnY1xZyvhPjClwtA6yEG2hJs890zwOjRwHN",
	"Ne4rshGm8w2YCD+ab2geMJeCc844XvF0ZHPZPta7weqM2rhUTJOFAOXm077CpN6ZPpPTNT82GL+f+GqO",
	"CMOGbc207SlAH9SRPxN444pACUmem7YEQ7TNz60jJTvoUVm6QWOaQNUrHKuONEjgSOQ58aG/gLg1/BDa",
	"FnbbepSIu6lhNLjAwwIocRfuMcbARfEfjCNpOcreN7VH+NELDIxH0HjJODS1RiMbRBrdEnBhUF4H+qlU",
	"Up0u99Zpp0BzPJ2IKTSlXbjmpqA6C4wkwTn6MYaXsamRNaA46gbN9QLKN3WJU8PdgUn3HGsrO0L2K16h",
	"TeVMqAyTtDo1sGKKwyhuX1WuvQH0xaBvEdnuWlIrOVfZiewpax9qxpRxAYpZHklLeVF/DOrAYf7bbIP/",
	"xq7pDs/AHWZdOV3Cn1xhxytbl21IPevQrH2i2OKaq9L0v8Vl6chAuEYx7v/BqJXw1mDvxrVVPHUVRDzy",
	"Fr6KJ7oUdaJ5m2dR0cXoEBRe3O4oDpdQHKNqHEjMedvcq6RW+9p43FB6TjqYTUa1SxXVlGyrLmLrG8Yg",
	"2LM/W1fRPnIQdcaHzvvscZ/53Ou9n93Qs8IQ9laC+oPkPkJ/8RklpKTMBZsbEYkmw0QZYK8EmWaBI4kt",
	"Iw8kNpNrJm3tJXt9KkUEOzyO38Ge5y2S2tsdHUtSSLhl0gZb6BVJ20802Hd6OA/kmEpBf557L0CLtgO0",
	"34fwjV7oE3dYnPVsH3GOJ8mb7qhPLEH8NY6+Nvlk2qBVltWNG1v1vw5WtLP3uKgmKyCUc4ES5aKShJJC",
	"ZJAT5Qqc5LCg6cZdvVRnPKWcZEwCVglhBVZWo0St6GIBEu/s2mKoPjaB0CKrVbE828U2Dsb32DZyFfpz",
	"XmbuC7FF9krmRHdpcaLbL+/Ww3ysC7upKAobGmiRP3pttb4Kh0EXRL+pBrgtcjiTlFtPpEchhBI8xBAp",
	"C7aknEMe7W3Pbj4ThxT0n2IA54Lx+KcuC1jCdMjQzLk9Qz+khx+pYzEeKUgryfQG86u8Z8L+Hs1L/6mW",
	"X1dLvj6ldoek9l0Td3zQSHvzFMVPwlZ3Loy7hK6DxtIzP6xpUebg9Oh338z+BI///CQ7ePzwT7M/Hzw9",
	"SOHJ02cHB/TZE/rw2eOH8OjPT58cwMP5t89mj7JHTx7Nnjx68u3TZ+njJw9nT7599qdv/DsQFtHmjYW/",
	"YS2H5OjNcXJqkG0WipbsL7Cx19ENd/p6GzRFzQ0FZfno0P/0v7ycGAEKnq5zv47cacxoqXWpDqfT1Wo1",
	"CbtMF1j+L9GiSpdTP06/0s+b4zqcb08+UJZsrNYIOu4XTOeYiYPf3v5wckqO3hxPGnUwOhwdTA4mD7H8",
	"Sgmclmx0OHqMPyHXL3Hdp0uguTaScTkeTQvQkqXK/eVU+MSVGjE/XTya+gjg9INLPbg0cBaxXDNfsqyO",
	"QPcvtY/tNmO82rpEWXB9S7lbXWMys1lVxFXJ4xnGiG3GjNn8avIcZ8HTmMGTC+PWy57vvqLHqmL1s2LV",
	"AWLPj9b3Coafnwle6POv8j3982Xk+O9952WRRwcHH+E1kXELiqfLLT9L8uQWUW/73jeeQBdcbxqvaG74",
	"Ceqn5+yEHn61EzrmeLPHKDBiFfTlePT0K16hY24EiuYEWwYJP30V+Ss/52LFfUuzOVdFQeUGt96gpEBo",
	"O10OquJ2qp27mzmsnyGo8BZc524dicw2ns/GRNWVnEvJhDEh8KHGDFIJFDd8IfEksakV5y6tgi1d/ero",
	"b3ju8Orob7YIY/QRu2B4W5C0rdx/Ah2pZfj9pnmIaaum/1zqc/zFvvv39eyFN92C7ipi3lXE/GorYn5M",
	"oyViZazrzFdKuOAJx4oFF0ACJ/Zjmh2f307YY2N/evD40w1/AvKCpUBOoSiFpJLlG/IrrzNmbmZo1HJT",
	"8SCHaasM9WqYN7ZCYKQEBYWmH4K/Epbtdh1b11SzViVrGn/fL6i14lLvxs3FRuM9YqaDP8tU4zrbkGf+",
	"yqtdj3HvguAkZooERxHfb/B9+53WR2tOwa2omAXSotfVnhH9qP7atd9e/KRa7HuaEZ9Q+UWoqycHTz4d",
	"BuEqvBaa/IhJWJ9faV5fScXZKlA2WLRr+sFfoNpDwbjLiW3V0n2wM6ZUjISOXR65K/ZbP+Fg9IlVhPZ+",
	"aF9rmBH21Rf9+5MxTdHcGftSdMSV3kO90wt3euHaeqHLUI1GsK+xTT9gAmqoDnoiiU/K/oHCxEG1OCkK",
	"X79IkDnodOleu+0cyQ29Yr5Vp2y76nZj/XL31vFN3jreI9B5R+BP85j013ziEOyWJCGv0RxCAfc5yX/E",
	"A4iPuSN/7Am9FhwIrJnCKpKWF+8OVWpzAS+FI1F8xf2wxHttOrgXF6cfmidQL5tzcHuFbmot/212hX0m",
	"ZHSrkeu7p12+gqddPr9XcSMJ6cxWQviOK7grpI20+CKU/cqM7VQR11wtK52JVZBY0hT7HZQk/6L3LUrS",
	"3bPid8+K3z0rfves+N2z4nfPit89K/51Pyv+9Z1Gd4N4H9HraZuwgSnTmHD27+mKMp3MhbTbU4LVvCIB",
	"1Pbo/02ZdjXknG+lhVEWYHZorAdmFY2DE9QWUWE+hnvEwT+nzYrIoasZ6kch94rXNkFQLYiZGKm4Zj7X",
	"GB/78fbclxf8vLNU7yzVO0v1zlK9s1TvLNU7S/WPZal+nmQHkiReUfvkzlhqJ7nL7fwD5XY2BnZtXqNB",
	"bsxhI99bD0E00Hzq6mfhebFQg9lUYS2u1AzHOClzikV519rfXMB6vN8+8ckQdVUZex3f6CDT4PEjcvLz",
	"0dOHj/7+6Om39QvW7bb3fP1QpTe5LcLb9hROgebPHe5WmYDS34ts01lXg94UMW2vaHNZmHEqIwWbIu8Y",
	"d2mgBRZtcxXIes7E5a0mSMQr2fbpuYuUA9Vco9y3bTl3FhF1l5Yd7H20qFlTT07iij19Vo1KECPHZo32",
	"+JdXn9dSV56MUTFCIRwbDsuqFPB1K8c/68Q0WgBPnJAnM5Ft/HMFrhJcS6XZEl3DGu2HNaSVkQzExDH1",
	"PXXfPSSIpQbDGEa0QGpQZRcQnsuz6mspWwxqq5K6/uK1C8ve+Ki+C27bW+7knpBkIUVV3rd16/kGndOi",
	"pHzjwy/GnsLKtPisI6YX3a5arOvy9ZTa/qVVQ5se7zt1f7dkISuqfF3VzBZWjReX6Zb/3E3xprjdrrIh",
	"dr7RQpwDZTf7i+hX2SU21iGnEmSi1zxSDq9T/O5fPqf3a9S/b6S4YMZVjKozG97VUfGe7FTDMlBAqIc7",
	"dw69Im5rx7d0Fd5g3FdDrhNns93YoFuCfe3JGziRC5pmc5KCZilVmITo6g9/ZGNPr48jnjaiiVex571L",
	"Wma33F3YHeHuZYoFoJtHhPAmrFI2C/uzGmZNpYQjl/PZosadlvijOLnfe+FThOJb/B3hDGqC76Gm6Eqv",
	"eVRLTZtHAaI5SoFA1M8a3eIJUA98+yAoeD/InkRAXhLqCrVhcFLLKtVnnGLQL3y3qX9I5EOZw4bRc98k",
	"HneOhIUdqDNO8aWNOhQYNZDmEKuQDeDtL1UtFqB0RxPPAc64a8V486pHwVIpEpupV4JEjT6xLQu6IXOa",
	"Y9T6d5CCzIzJHl58xVCZ0izP3amUGYaI+RnHcnhG6b9ixjwz4Hw0pT5pdbXowzfH+yHpbiG7fhEuxdTP",
	"VC399H1EBAM39rM9ePn0D8m0y+BFMT9+4QorHL/Ae8bNgVQP9092oFIwnkSZzOz47ly3y1vknnvWCBno",
	"fnO05Vb9jBvTWAv7InjzpujV2KEb+O7JopWO7WUBW/FxP9ePVSLw4uEO++AG+opE1NXdzv0HKj3Qefeu",
	"XnhjxPbWfmBfvoVKR192eaOdiS53xYTuigndFRPas5jQHhHQu9W9KxX1FZeKuisH+QXfXPyYptvHns2X",
	"XoRqstVCnH7Q633KwoRQWWaf65SQ2pFrBR42axWQ6Z8BMj0h5BTf4qRmD4ALkDTHJ5iVv87OFCnYYqmJ",
	"qtIUIDs840kLE1vp2wx8r/mvdXPPqoODx0AO7pN2Fxu2CBRvvytaqvjJPhLzHTkbnY26gCQU4gJcMQls",
	"nVV4LGs77YT6bw7sGf9F9hauoBsbWlnSsgSzqalqPmcpswTPhXEFFqKTz8YFfgFpkAOjTxVh2tbtQmpi",
	"HqDLOqHuDZyYyd3f3a9QOfqowyzxVHLDdlesI/qf+xQR/Vcxr1+ApixXdYZ7xJtCv6bLWSuqGsGtdcrY",
	"J0Yr/5s7fHaj5OwcwpxTPOhfUZn5FpH3h2z9Jf9qXeR1eFekJoO1NwK6iM7r0VjzgHz9Jn88KToXChKL",
	"nIo9loIfjALAECjFCCh1Dwz7NzQNDCND1GAn8eaGTSAfHpPxRWKrr0ciw/a7q85eh8A6AecIXL88g1mk",
	"9Yr4V/OZ6hExXOQ5cRe44wMa9ZQMPNp33E+i7Y50ztJzyIhhSP+K84CtSO7VpcHwVdbVcuNvC1h9d39C",
	"yBG376j7B1rbIc3O4PwbvW38daih26ovktiVArsAeUMu8mC2844Cw2I3HMoC2T6QXvMBBqKriOe0b62Y",
	"iKPUcVsCprJY7OOhfP12R7fP9Q2PLqTbszw+u+1xlxTzSQvdhQkKrUJ3N/BQ6sdMYhaIRcK/r4PGYv2y",
	"zrv3xiRSIC+8Hdk8F3M4nWLt2aVQejoyVl77KZnwo1EndGEhODutlOwC61a9v/z/AQAA//8kewwmz9gA",
	"AA==",
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
