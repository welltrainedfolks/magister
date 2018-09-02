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

package admin

import (
	// stdlib
	"net/http"

	// local
	"github.com/welltrainedfolks/magister/internal/templater"

	// other
	"github.com/labstack/echo"
)

// adminGET is a handler for admin interface.
func adminGET(ec echo.Context) error {
	if !ec.Get("AUTHORIZED").(bool) {
		return ec.Redirect(http.StatusMovedPermanently, "/login_required/")
	}

	data := make(map[string]string)

	var tabTpl string
	tab := ec.Param("tab")
	if tab == "index" {
		tabTpl = templater.GetRawTemplate(ec, "admin/index.html", nil)
	} else if tab == "packages" {
		tabTpl = templater.GetRawTemplate(ec, "admin/packages.html", nil)
	}

	data["tab.data"] = tabTpl

	// Tabs.
	data["tab.index.active"] = ""
	data["tab.packages.active"] = ""
	// ...and activate required.
	data["tab."+tab+".active"] = "is-active"

	htmlData := templater.GetTemplate(ec, "admin/skeleton.html", data)

	return ec.HTML(http.StatusOK, htmlData)
}
