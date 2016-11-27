package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"regexp"
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
	target        *url.URL
	proxy         *httputil.ReverseProxy
	routePatterns []*regexp.Regexp
}

// New Proxy.
func New(target string) *Proxy {
	url, _ := url.Parse(target)

	return &Proxy{target: url, proxy: httputil.NewSingleHostReverseProxy(url)}
}

func (p *Proxy) handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-GoProxy", "GoProxy")

	if p.routePatterns == nil || p.parseWhiteList(r) {
		p.proxy.ServeHTTP(w, r)
	}
}

func (p *Proxy) parseWhiteList(r *http.Request) bool {
	for _, regexp := range p.routePatterns {
		fmt.Println(r.URL.Path)
		if regexp.MatchString(r.URL.Path) {
			return true
		}
	}
	fmt.Println("Not accepted routes ", r.URL.Path)
	return false
}

func main() {
	const (
		defaultPort             = ":8080"
		defaultPortUsage        = "default server port, ':80', ':8080'..."
		defaultTarget           = "http://localhost:9000"
		defaultTargetUsage      = "default redirect url, 'localhost:9000'"
		defaultWhiteRoutes      = `.*`
		defaultWhiteRoutesUsage = "list of white route as regexp, '/path1*,/path2*...."
	)

	// flags
	port := flag.String("port", defaultPort, defaultPortUsage)
	url := flag.String("url", defaultTarget, defaultTargetUsage)
	routesRegexp := flag.String("routes", defaultWhiteRoutes, defaultWhiteRoutesUsage)

	flag.Parse()

	fmt.Println("server will run on: ", *port)
	fmt.Println("redirecting to: ", *url)
	fmt.Println("accepted routes: ", *routesRegexp)

	//
	reg, _ := regexp.Compile(*routesRegexp)
	routes := []*regexp.Regexp{reg}

	// proxy
	proxy := New(*url)
	proxy.routePatterns = routes

	// server
	http.HandleFunc("/posts", proxy.handle)
	http.ListenAndServe(*port, nil)
}
