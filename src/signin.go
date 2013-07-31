// Copyright 2013 GCP Lab. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package badge

import (
	"html/template"
	"net/http"
)

func init() {
	http.HandleFunc("/signin", signinHandler)
}

type UserInfo struct {
	Name  string
	ID    string
	Pwd   string
	Email string
}

func signinHandler(w http.ResponseWriter, r *http.Request) {
	signin := UserInfo{Name: "Admin", ID: "admin", Pwd: "gran2t", Email: "admin@gmail.com"}
	temp, _ := template.ParseFiles("templates/signin.html")
	temp.Execute(w, signin)
}
