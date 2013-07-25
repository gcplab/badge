// Copyright 2013 GCP Lab. All rights reserved.
// Use of this sorce code is governed by the MIT license
// that can be found in the LICENSE file.

package badge

import (
	"net/http"
	"text/template"
)

var (
	eventsTmpl *template.Template =
		template.Must(template.ParseFiles("templates/events.html"))
)

func init() {
	http.HandleFunc("/events", eventsHandler)
}

func eventsHandler(w http.ResponseWriter, r *http.Request) {
	// Do nothing now...
	err := eventsTmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
