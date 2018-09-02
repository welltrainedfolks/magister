// Copyright (c) 2018, Well Trained Folks and contributors.
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject
// to the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
// CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
// TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE
// OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package http

import (
	// local
	"github.com/welltrainedfolks/magister/assets/compiled"
	"github.com/welltrainedfolks/magister/internal/config"

	// other
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/rs/zerolog/log"
)

var (
	// E is a HTTP server.
	E *echo.Echo
)

// Initialize initializes package.
func Initialize() {
	log.Info().Msg("Initializing HTTP server...")

	authRequiredEndpoints = []string{}

	E = echo.New()
	E.Use(echoReqLogger())
	E.Use(middleware.Recover())
	E.Use(loginStateChecker())
	E.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "form:_magcsrf",
		ContextKey:  "CSRFTOKEN",
	}))
	E.DisableHTTP2 = true
	E.HideBanner = true
	E.HidePort = true

	// Static files.
	E.GET("/static/*", echo.WrapHandler(assets.Handler))

	// Index.
	E.GET("/", indexGET)

	// Default handler for 404 and invalid method.
	echo.NotFoundHandler = NotFoundGET
	echo.MethodNotAllowedHandler = NotFoundGET
}

// Shutdown shutdowns HTTP server.
func Shutdown() {
	log.Info().Msg("Shutting down HTTP server gracefully...")
	err := E.Shutdown(nil)
	if err != nil {
		log.Error().Msgf("Failed to shutdown HTTP server gracefully: %s", err.Error())
	}
}

// StartListening starts echo's HTTP server.
func StartListening() {
	listenAddress := config.Config.HTTP.Address + ":" + config.Config.HTTP.Port
	go func() {
		E.Start(listenAddress)
	}()

	log.Info().Msgf("Starting to listening for HTTP requests on %s", listenAddress)
}
