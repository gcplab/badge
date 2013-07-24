// Copyright 2013 GCP Lab. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package badge

import (
  "net/http"
  "html/template"
)

func init() {
  http.HandleFunc("/admin", admin)
}

type Admin struct {
  Name string
  Grant string
}

func admin(w http.ResponseWriter, r *http.Request) {
  admin := Admin{Name: "admin", Grant: "gran2t"}
  temp, _ := template.ParseFiles("templates/admin.html")
  temp.Execute(w, admin)
}
