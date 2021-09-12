// Copyright (c) 2019, Boise State University All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package favicon

import (
	_ "embed"
	"net/http"
)

//go:embed testdata/favicon.ico
var icon []byte

func Favicon(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/vnd.microsoft.icon")
	w.Write(icon)
}
