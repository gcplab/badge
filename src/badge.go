// Copyright 2013 GCP Lab. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package badge

import (
	"net/http"
	"text/template"
)

var (
	indexTmpl *template.Template = template.Must(template.ParseFiles("templates/index.html"))
)

func init() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/login", loginHandler)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	err := indexTmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type User struct {
	Id  string
	Pwd string
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PostFormValue("id")
	pwd := r.PostFormValue("pwd")
	// TODO id, pwd가 ""면 에러 리턴

	user := User{Id: id, Pwd: pwd}

	err := indexTmpl.Execute(w, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
