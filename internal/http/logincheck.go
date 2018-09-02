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
	// stdlib
	"errors"
	"net/http"
	"strings"

	// local
	"github.com/welltrainedfolks/magister/internal/sessionkeys"

	// other
	"github.com/labstack/echo"
	"github.com/rs/zerolog/log"
)

// Checks every request's cookies to determine if user is logged in.
func loginStateCheck(ec echo.Context, next echo.HandlerFunc) error {
	ec.Set("AUTHORIZED", false)
	// Check for session cookie.
	log.Debug().Msgf("Cookies: %+v", ec.Cookies())
	sessionkey, err1 := ec.Cookie("s3ss1onk3y")
	if err1 != nil {
		log.Debug().Msgf("No s3ss1onk3y cookie, user isn't logged in")
	}

	// Check cookie validity.
	if sessionkey != nil {
		uid, authed := sessionkeys.CheckSessionKey(sessionkey.Value)
		if authed {
			log.Debug().Msg("User is authorized")
			ec.Set("AUTHORIZED", true)
			ec.Set("UID", uid)
		}
	}

	if !ec.Get("AUTHORIZED").(bool) {
		// If user isn't authorized - check if current URL requires auth.
		var authRequired bool
		for _, ep := range authRequiredEndpoints {
			if strings.Contains(ec.Request().RequestURI, ep) {
				authRequired = true
			}
		}

		if authRequired {
			log.Debug().Msg("Trying to access endpoint which require authorization without it!")
			ec.Redirect(http.StatusMovedPermanently, "/login_required/")
			return errors.New("User isn't logged in")
		}
	}

	next(ec)
	return nil
}

// Wrapper around previous function.
func loginStateChecker() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return loginStateCheck(c, next)
		}
	}
}
