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

package templater

import (
	// stdlib
	"errors"
	"net/http"
	"strconv"
	"strings"

	// local
	"github.com/welltrainedfolks/magister/assets/compiled"
	"github.com/welltrainedfolks/magister/common"
	"github.com/welltrainedfolks/magister/internal/config"

	// other
	"github.com/labstack/echo"
	"github.com/rs/zerolog/log"
)

var (
	// Actions which is called on every request.
	// Returned data will be replaced in template.
	// They're registered with RegisterTemplateName function defined
	// below.
	// Currently these strings are registered:
	// - {user.name}
	actions map[string]func(ec echo.Context) string
)

// GetErrorFlash returns formatted HTML for errors flash.
func GetErrorFlash(ec echo.Context, errors []string) string {
	errorsData := make(map[string]string)

	// Format errors.
	if len(errors) != 0 {
		errorsString := "<ul>"
		for _, err := range errors {
			errorsString += "<li>" + err + "</li>"
		}
		errorsString += "</ul>"
		errorsData["errors"] = errorsString

		return GetRawTemplate(ec, "partials/error.html", errorsData)
	}

	return ""
}

// GetErrorTemplate returns formatted error template.
// If error.html wasn't found - it will return "error.html not found"
// message as simple string.
func GetErrorTemplate(ec echo.Context, errorText string) string {
	// Getting main error template.
	mainhtml := GetTemplate(ec, "error.html", map[string]string{"error": errorText})

	return mainhtml
}

// GetRawTemplate returns only raw template data.
func GetRawTemplate(ec echo.Context, templateName string, data map[string]string) string {
	// Getting main template.
	tplRaw, err := assets.ReadFile(templateName)
	if err != nil {
		ec.String(http.StatusBadRequest, templateName+" not found.")
		return ""
	}

	tpl := string(tplRaw)

	// Replace data with data returned by actions functions.
	for key, handler := range actions {
		tpl = strings.Replace(tpl, "{"+key+"}", handler(ec), -1)
	}

	// Replace placeholders with data from data map.
	for placeholder, value := range data {
		tpl = strings.Replace(tpl, "{"+placeholder+"}", value, -1)
	}

	// CSRF.
	tpl = strings.Replace(tpl, "{csrf_token}", ec.Get("CSRFTOKEN").(string), -1)

	return tpl
}

// GetSuccessFlash is identical to GetErrorFlash, but returns successes.
func GetSuccessFlash(ec echo.Context, successes []string) string {
	successData := make(map[string]string)

	// Format errors.
	if len(successes) != 0 {
		successesString := "<ul>"
		for _, err := range successes {
			successesString += "<li>" + err + "</li>"
		}
		successesString += "</ul>"
		successData["successes"] = successesString

		return GetRawTemplate(ec, "partials/success.html", successData)
	}

	return ""
}

// GetTemplate returns formatted template that can be outputted to client.
func GetTemplate(ec echo.Context, name string, data map[string]string) string {
	log.Debug().Msgf("Requested template '%s'", name)

	// Getting main template.
	mainhtml, err := assets.ReadFile("main.html")
	if err != nil {
		ec.String(http.StatusBadRequest, "main.html not found.")
		return ""
	}

	// Getting navigation.
	navhtml, err1 := assets.ReadFile("navigation.html")
	if err1 != nil {
		ec.String(http.StatusBadRequest, "navigation.html not found.")
		return ""
	}

	// ToDo: logged in or not checks.
	var err4 error
	var loginBarHTML []byte
	if !ec.Get("AUTHORIZED").(bool) {
		loginBarHTML, err4 = assets.ReadFile("loginbar-notloggedin.html")
	} else {
		loginBarHTML, err4 = assets.ReadFile("loginbar-loggedin.html")
	}

	if err4 != nil {
		ec.String(http.StatusBadRequest, "loginbar-*.html not found.")
		return ""
	}

	// Getting footer.
	footerhtml, err2 := assets.ReadFile("footer.html")
	if err2 != nil {
		ec.String(http.StatusBadRequest, "footer.html not found.")
		return ""
	}

	// Format main template.
	tpl := strings.Replace(string(mainhtml), "{navigation}", string(navhtml), 1)
	tpl = strings.Replace(tpl, "{footer}", string(footerhtml), 1)
	// Version.
	tpl = strings.Replace(tpl, "{code.version}", common.VERSION, 1)
	tpl = strings.Replace(tpl, "{code.build}", strconv.Itoa(common.BUILD), 1)
	tpl = strings.Replace(tpl, "{code.build_date}", common.BUILDDATE, 1)

	// Get requested template.
	reqhtml, err3 := assets.ReadFile(name)
	if err3 != nil {
		ec.String(http.StatusBadRequest, name+" not found.")
		return ""
	}

	// Replace basic things.
	tpl = strings.Replace(tpl, "{site.name}", config.Config.Site.Name, -1)
	tpl = strings.Replace(tpl, "{loginBar}", string(loginBarHTML), -1)

	// Replace documentBody.
	tpl = strings.Replace(tpl, "{documentBody}", string(reqhtml), 1)

	// Replace data with data returned by actions functions.
	for key, handler := range actions {
		tpl = strings.Replace(tpl, "{"+key+"}", handler(ec), -1)
	}

	// Replace placeholders with data from data map.
	for placeholder, value := range data {
		tpl = strings.Replace(tpl, "{"+placeholder+"}", value, -1)
	}

	// CSRF.
	tpl = strings.Replace(tpl, "{csrf_token}", ec.Get("CSRFTOKEN").(string), -1)

	return tpl
}

func GetTextTemplate(templateName string, data map[string]string) string {
	log.Debug().Msgf("Requested text template: '%s'", templateName)

	tplRaw, err := assets.ReadFile(templateName)
	if err != nil {
		log.Error().Msgf("Failed to get text template '%s': %s", templateName, err.Error())
		return ""
	}

	// Replace basic variables.
	tpl := strings.Replace(string(tplRaw), "{site.name}", config.Config.Site.Name, -1)
	tpl = strings.Replace(tpl, "{code.version}", common.VERSION, 1)
	tpl = strings.Replace(tpl, "{code.build}", strconv.Itoa(common.BUILD), 1)
	tpl = strings.Replace(tpl, "{code.build_date}", common.BUILDDATE, 1)

	// Text templates have no possibility to access echo's Context
	// because they're obtained asynchronously compared to general HTTP
	// request. So there will be no placeholders replacement based on
	// actions-returned data.

	// Replace placeholders with data from data map.
	for placeholder, value := range data {
		tpl = strings.Replace(tpl, "{"+placeholder+"}", value, -1)
	}

	return tpl
}

func Initialize() {
	actions = make(map[string]func(ec echo.Context) string)
}

func RegisterTemplateName(name string, handler func(ec echo.Context) string) error {
	_, alreadyRegistered := actions[name]
	if alreadyRegistered {
		return errors.New("Template name '" + name + "' already registered")
	}

	actions[name] = handler
	return nil
}
