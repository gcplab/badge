// Copyright 2013 GCP Lab. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package badge

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", handle)
}

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<html><body>Hello, World! 세상아 안녕!</body></html>")
}
