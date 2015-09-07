package main

import (
	"crypto/tls"
	"net/http"
)

var client *http.Client

func init() {
	tr := &http.Transport{
		DisableCompression: true,
	}

	client = &http.Client{Transport: tr}
}
