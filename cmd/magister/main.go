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

package main

import (
	// stdlib
	"flag"
	"os"
	"os/signal"
	"syscall"

	// local
	"github.com/welltrainedfolks/magister/admin"
	"github.com/welltrainedfolks/magister/common"
	"github.com/welltrainedfolks/magister/internal/config"
	"github.com/welltrainedfolks/magister/internal/database"
	"github.com/welltrainedfolks/magister/internal/http"
	"github.com/welltrainedfolks/magister/internal/mailsender"
	"github.com/welltrainedfolks/magister/internal/templater"
	"github.com/welltrainedfolks/magister/users"

	// other
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Info().Msgf("Starting MAGISTER, version %s (build %d, built on %s from revision %s, branch %s)", common.VERSION, common.BUILD, common.BUILDDATE, common.REVISION, common.BRANCH)

	// Initialize everything that wants CLI flag(s).
	config.Initialize()

	// Parse CLI flags.
	flag.Parse()

	// Post-flags-parsing initialization.
	config.LoadConfiguration()
	templater.Initialize()
	database.Initialize()
	http.Initialize()

	// Initialize modules.
	admin.Initialize()
	mailsender.Initialize()
	users.Initialize()

	// Start HTTP server.
	http.StartListening()

	// CTRL+C handler.
	signalHandler := make(chan os.Signal, 1)
	shutdownDone := make(chan bool, 1)
	signal.Notify(signalHandler, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-signalHandler
		log.Info().Msg("Starting MAGISTER shutdown...")

		http.Shutdown()

		shutdownDone <- true
	}()

	<-shutdownDone
	os.Exit(0)
}
