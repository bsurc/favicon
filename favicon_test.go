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
