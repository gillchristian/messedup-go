package services

import "net/url"

// Service implements reference to a service
type Service struct {
	name   string
	target *url.URL
	auth   bool
}

// Services: an in memory Service collection.
type Services []Service
