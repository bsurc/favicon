// Copyright (c) 2019, Boise State University All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestFavicon(t *testing.T) {
	want, _ := ioutil.ReadFile("./testdata/favicon.ico")
	if got := bytes.Compare(icon, want); got != 0 {
		t.Errorf("got: %d, want: 0", got)
	}
}
