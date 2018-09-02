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

package users

import (
	// local
	"github.com/welltrainedfolks/magister/internal/http"
	"github.com/welltrainedfolks/magister/internal/templater"

	// other
	"github.com/rs/zerolog/log"
)

func Initialize() {
	log.Info().Msg("Initializing 'users' module...")

	// Template actions.
	templater.RegisterTemplateName("user.name", GetCurrentlyLoggedInUserName)

	// Login.
	http.E.GET("/login/", loginGET)
	http.E.POST("/login/", loginPOST)

	// Login required form.
	http.E.GET("/login_required/", loginRequiredGET)

	// Already logged in message.
	http.E.GET("/already_logged_in/", alreadyLoggedInGET)
	// Already logged out message.
	http.E.GET("/already_logged_out/", alreadyLoggedOutGET)

	// Logout.
	http.E.GET("/logout/", logoutGET)

	// Profile.
	http.E.GET("/profile/:tab/", profileGET)
	http.E.POST("/profile/:tab/", profilePOST)
}
