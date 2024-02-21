package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	r := chi.NewRouter()

	proxy := NewReverseProxy("hugo", "1313")

	r.Use(proxy.ReverseProxy)
	// ...
	r.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
	})

	go TimeWorker()
	go GraphWorker()
	go BinaryWorker()

	http.ListenAndServe(":8080", r)
}

type ReverseProxy struct {
	host  string
	port  string
	proxy *httputil.ReverseProxy
}

func NewReverseProxy(host, port string) *ReverseProxy {
	targetURL, _ := url.Parse(fmt.Sprintf("http://%s:%s", host, port))
	return &ReverseProxy{
		host:  host,
		port:  port,
		proxy: httputil.NewSingleHostReverseProxy(targetURL),
	}
}

func (rp *ReverseProxy) ReverseProxy(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/" {
			w.Write([]byte("Hello from API"))
		} else {
			rp.proxy.ServeHTTP(w, r)
		}
	})
}
