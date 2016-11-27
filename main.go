package main

import (
	"net/http/httputil"
	"net/url"
	// "flag"
	// "fmt"
	// "regexp"
	//
	// "net/http"
	// "net/http/httputil"
	// "net/url"
	// "github.com/gorilla/mux"
)

// Proxy struct.
type Proxy struct {
	name   string
	target *url.URL
	proxy  *httputil.ReverseProxy
}
