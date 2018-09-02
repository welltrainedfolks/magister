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
	//"strings"
	//"time"

	// local
	h "github.com/welltrainedfolks/magister/internal/http"
	"github.com/welltrainedfolks/magister/internal/templater"

	// other
	"github.com/labstack/echo"
	"github.com/rs/zerolog/log"
)

type PasswordChangeRequest struct {
	Login               string `form:"login"`
	CurrentPassword     string `form:"current-password"`
	NewPassword         string `form:"new-password"`
	NewPasswordRepeated string `form:"new-repeated-password"`
}

func profileGET(ec echo.Context) error {
	if !ec.Get("AUTHORIZED").(bool) {
		return ec.Redirect(http.StatusMovedPermanently, "/login_required/")
	}

	// This data should be on every profile tab.
	data := make(map[string]string)
	data["csrf_token"] = ec.Get("CSRFTOKEN").(string)

	// Get currently active tab and obtain tab-specific data.
	var tabTpl string
	tab := ec.Param("tab")
	if tab == "general" {
		u := GetCurrentlyLoggedInUser(ec)
		data["user.login"] = u.Email
		data["user.registered"] = u.CreatedAt.Format("2006-01-02 15:04:05")
		tabTpl = templater.GetRawTemplate(ec, "profile/general.html", data)
	} else if tab == "password" {
		tabTpl = templater.GetRawTemplate(ec, "profile/password.html", map[string]string{"errorsDiv": "", "successDiv": "", "csrf_token": ec.Get("CSRFTOKEN").(string)})
	}

	data["tab.data"] = tabTpl
	// Populate tabs list.
	data["tab.general.active"] = ""
	data["tab.contacts.active"] = ""
	data["tab.password.active"] = ""
	data["tab.forums-general.active"] = ""
	// Set active.
	data["tab."+tab+".active"] = "is-active"

	htmlData := templater.GetTemplate(ec, "profile/skeleton.html", data)

	return ec.HTML(http.StatusOK, htmlData)
}

func profilePOST(ec echo.Context) error {
	if !ec.Get("AUTHORIZED").(bool) {
		return ec.Redirect(http.StatusMovedPermanently, "/login_required/")
	}

	tab := ec.Param("tab")
	log.Debug().Msgf("Profile POST on tab %s", tab)

	if tab == "password" {
		return profilePasswordPOST(ec)
	}

	return h.NotFoundGET(ec)
}

func profilePasswordPOST(ec echo.Context) error {
	req := &PasswordChangeRequest{}
	if err := ec.Bind(req); err != nil {
		log.Error().Msgf("Failed to read form data: %s", err.Error())
		// ToDo: send registration error mail.
	}

	var errors []string

	u := GetUser(req.Login)
	if u == nil {
		errors = append(errors, "General system error, please try again later")
	}

	if !u.CheckPassword(req.CurrentPassword) {
		errors = append(errors, "Invalid current password entered")
	}

	if req.NewPassword != req.NewPasswordRepeated {
		errors = append(errors, "New password and new password confirmation doesn't match")
	}

	if len(errors) != 0 {
		tabTpl := templater.GetRawTemplate(ec, "profile/password.html", map[string]string{"errorsDiv": templater.GetErrorFlash(ec, errors), "successDiv": "", "csrf_token": ec.Get("CSRFTOKEN").(string)})

		formData := map[string]string{
			"tab.data":            tabTpl,
			"tab.password.active": "is-active",
		}

		profileTpl := templater.GetTemplate(ec, "profile/skeleton.html", formData)

		return ec.HTML(http.StatusBadRequest, profileTpl)
	}

	u.CreatePassword(req.NewPassword)
	u.Save()

	tabTpl := templater.GetRawTemplate(ec, "profile/password.html", map[string]string{"errorsDiv": templater.GetErrorFlash(ec, []string{}), "successDiv": templater.GetSuccessFlash(ec, []string{"Password successfully changed"}), "csrf_token": ec.Get("CSRFTOKEN").(string)})

	formData := map[string]string{
		"tab.data":            tabTpl,
		"tab.password.active": "is-active",
	}

	profileTpl := templater.GetTemplate(ec, "profile/skeleton.html", formData)

	return ec.HTML(http.StatusBadRequest, profileTpl)
}
