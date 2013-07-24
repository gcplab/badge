// Copyright 2013 GCP Lab. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package badge

import (
	"net/http"
	"text/template"
)

var (
	loginTmpl *template.Template = template.Must(template.ParseFiles("templates/login.html"))
)

func init() {
    http.HandleFunc("/login", loginHandler)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	err := loginTmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
