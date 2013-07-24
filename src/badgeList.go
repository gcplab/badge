// Copyright 2013 GCP Lab. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package badge

import (
	"net/http"
	"text/template"
	"log"
)

var (
	badgeListTmpl *template.Template = template.Must(template.ParseFiles("templates/badgeList.html"))
)

func init() {
	http.HandleFunc("/badgeList", badgeListHandler)
	log.Println("TEST!!!");
}

func badgeListHandler(w http.ResponseWriter, r *http.Request) {
	err := badgeListTmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
