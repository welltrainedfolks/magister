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

	// local
	"github.com/welltrainedfolks/magister/common"
	"github.com/welltrainedfolks/magister/internal/config"
	"github.com/welltrainedfolks/magister/internal/database"
	"github.com/welltrainedfolks/magister/internal/http"
	"github.com/welltrainedfolks/magister/internal/templater"
	"github.com/welltrainedfolks/magister/users"

	// other
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	// CLI flags.

	// Users-related actions.
	userEmail    string
	userName     string
	userPassword string

	// Users registration.
	actionUserDeletion     bool
	actionUserRegistration bool
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Info().Msgf("Starting magisterctl, version %s (build %d, built on %s from revision %s, branch %s)", common.VERSION, common.BUILD, common.BUILDDATE, common.REVISION, common.BRANCH)

	// Initialize main CLI flags.
	flag.StringVar(&userEmail, "user_email", "", "E-Mail address for user.")
	flag.StringVar(&userName, "user_name", "", "User's name.")
	flag.StringVar(&userPassword, "user_password", "", "User's password.")
	flag.BoolVar(&actionUserDeletion, "user_delete", false, "Deletes user. Require \"user_name\" parameter.")
	flag.BoolVar(&actionUserRegistration, "user_register", false, "Register user. Require all \"user_*\" variables.")

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
	users.Initialize()

	if actionUserDeletion {
		deleteUser()
	} else if actionUserRegistration {
		registerUser()
	}
}

func deleteUser() {
	if userName == "" {
		log.Error().Msg("User's login wasn't provided")
		flag.PrintDefaults()
	}

	user := users.GetUserByLogin(userName)
	if user == nil {
		log.Fatal().Msgf("User '%s' wasn't found")
	}

	err := user.Delete()
	if err != nil {
		log.Error().Msgf("Failed to delete user: %s", err.Error())
	} else {
		log.Info().Msg("User successfully deleted")
	}
}

func registerUser() {
	var failed bool

	if userName == "" {
		log.Error().Msg("User's login wasn't provided")
		failed = true
	}

	if userEmail == "" {
		log.Error().Msg("User's E-Mail address wasn't provided")
		failed = true
	}

	if userPassword == "" {
		log.Error().Msg("User's password wasn't provided")
		failed = true
	}

	if failed {
		flag.PrintDefaults()
	}

	log.Debug().Msgf("Trying to register user:\n\tLogin: %s\n\tPassword: %s\n\tE-Mail: %s", userName, userPassword, userEmail)

	// Check if user already exist.
	u := users.GetUserByLogin(userName)
	if u != nil {
		log.Fatal().Msgf("User with login '%s' already registered!", userName)
	}

	user := users.NewUser(userName, userEmail, userPassword)
	user.SetActive()
	log.Info().Msgf("Registered new user: %+v", user)
}
