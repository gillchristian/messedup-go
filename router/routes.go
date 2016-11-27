package main

// Route implements the route interface
type Route struct {
	name    string
	method  string
	pattern string
	handler http.handlerFunc
}
