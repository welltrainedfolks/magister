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
	// stdlib
	"net/http"
	"strings"
	"time"

	// local
	"github.com/welltrainedfolks/magister/internal/config"
	"github.com/welltrainedfolks/magister/internal/sessionkeys"
	"github.com/welltrainedfolks/magister/internal/templater"

	// other
	"github.com/labstack/echo"
	"github.com/rs/zerolog/log"
)

type LoginRequest struct {
	Login    string `form:"login"`
	Password string `form:"password"`
}

func loginGET(ec echo.Context) error {
	if ec.Get("AUTHORIZED").(bool) {
		return ec.Redirect(http.StatusMovedPermanently, "/already_logged_in/")
	}

	htmlData := templater.GetTemplate(ec, "users/login.html", map[string]string{
		"login":     "",
		"errorsDiv": "",
	})

	return ec.HTML(http.StatusOK, htmlData)
}

func loginPOST(ec echo.Context) error {
	log.Debug().Msg("Login request")

	if ec.Get("AUTHORIZED").(bool) {
		log.Warn().Msgf("User #%d tried to access /login/ POST while authorized!", ec.Get("UID").(int))
		return ec.Redirect(http.StatusMovedPermanently, "/invalid_request/")
	}

	req := &LoginRequest{}
	if err := ec.Bind(req); err != nil {
		log.Error().Msgf("Failed to read form data: %s", err.Error())
		// ToDo: send registration error mail.
	}

	// Check for errors.
	var errors []string

	if req.Login == "" {
		errors = append(errors, "Login should not be empty.")
	}

	if req.Password == "" {
		errors = append(errors, "Password should not be empty.")
	}

	var u *User
	if req.Login != "" {
		var userFound bool
		u = GetUser(req.Login)
		if u != nil {
			userFound = true
		}

		if !userFound {
			u = GetUserByLogin(req.Login)
			if u != nil {
				userFound = true
			}
		}

		if !userFound {
			log.Error().Msgf("User with login or email '%s' wasn't found", req.Login)
			errors = append(errors, "Invalid login or password.")
		}
	}

	var passwordValid bool
	if u != nil && !u.CheckPassword(req.Password) {
		log.Error().Msgf("Invalid password for user '%s'", req.Login)
		errors = append(errors, "Invalid login or password.")
	} else if u != nil && u.CheckPassword(req.Password) {
		passwordValid = true
	}

	if len(errors) != 0 {
		log.Warn().Msg("Some errors were detected, showing login dialog with error messages")
		registerData := map[string]string{
			"login":     req.Login,
			"errorsDiv": templater.GetErrorFlash(ec, errors),
		}
		registerTpl := templater.GetTemplate(ec, "users/login.html", registerData)

		return ec.HTML(http.StatusBadRequest, registerTpl)
	}

	// Check provided data.
	var tpl string
	if u != nil && passwordValid {
		log.Debug().Msg("User exists and password is valid")
		// Check if user is activated.
		if !u.IsActive {
			tpl = templater.GetTemplate(ec, "users/login_failed_not_activated.html", nil)
			return ec.HTML(http.StatusOK, tpl)
		}
		log.Debug().Msg("Valid credentials, sending authorization cookies")

		key := sessionkeys.NewSessionKey(u.ID)

		cookieKey := new(http.Cookie)
		cookieKey.Name = "s3ss1onk3y"
		cookieKey.Value = key
		cookieKey.Expires = time.Now().UTC().Add(time.Hour * time.Duration(24*config.Config.HTTP.SessionValidityDays))
		cookieKey.Domain = strings.Split(strings.Split(config.Config.HTTP.Domain, "/")[2], ":")[0]
		cookieKey.Path = "/"
		log.Debug().Msgf("Cookie prepared: %+v", cookieKey)
		ec.SetCookie(cookieKey)

		tpl = templater.GetTemplate(ec, "users/login_success.html", nil)
	} else {
		log.Error().Msg("Invalid credentials passed")
	}

	return ec.HTML(http.StatusOK, tpl)
}
