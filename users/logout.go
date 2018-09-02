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
	"time"

	// local
	"github.com/welltrainedfolks/magister/internal/sessionkeys"
	"github.com/welltrainedfolks/magister/internal/templater"

	// other
	"github.com/labstack/echo"
	"github.com/rs/zerolog/log"
)

func logoutGET(ec echo.Context) error {
	if !ec.Get("AUTHORIZED").(bool) {
		return ec.Redirect(http.StatusMovedPermanently, "/already_logged_out/")
	}
	log.Debug().Msgf("Logging out user #%d", ec.Get("UID").(int))

	// ToDo: check for error and do something.
	c, err := ec.Cookie("s3ss1onk3y")
	if c.Expires == time.Unix(0, 0) || err != nil {
		return ec.HTML(http.StatusOK, templater.GetTemplate(ec, "users/already_logged_out.html", map[string]string{}))
	}

	c.Expires = time.Unix(0, 0)
	ec.SetCookie(c)
	sessionkeys.DeleteSessionKey(c.Value)

	return ec.HTML(http.StatusOK, templater.GetTemplate(ec, "users/logged_out.html", map[string]string{}))
}
