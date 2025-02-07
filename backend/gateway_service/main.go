package main

import (
	"diplomaPorject/backend/gateway_service/middleware"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func main() {
	r := mux.NewRouter()

	r.PathPrefix("/blogs").Handler(middlewares.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		serveReverseProxy(w, r)
	})))

	r.PathPrefix("/login").Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		serveReverseProxy(w, r)
	}))
	r.PathPrefix("/register").Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		serveReverseProxy(w, r)
	}))
	r.PathPrefix("/profile").Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		serveReverseProxy(w, r)
	}))

	r.PathPrefix("/admin/events").Handler(middlewares.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		serveReverseProxy(w, r)
	})))

	r.PathPrefix("/events").Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		serveReverseProxy(w, r)
	}))

	fmt.Println("Gateway running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func serveReverseProxy(w http.ResponseWriter, r *http.Request) {
	blogServiceURL := "http://localhost:8081"
	authServiceURL := "http://localhost:8082"
	eventsServiceURL := "http://localhost:8083"

	var target string
	switch {
	case strings.HasPrefix(r.URL.Path, "/blogs"):
		target = blogServiceURL
	case strings.HasPrefix(r.URL.Path, "/login") || strings.HasPrefix(r.URL.Path, "/register") || strings.HasPrefix(r.URL.Path, "/profile"):
		target = authServiceURL
	case strings.HasPrefix(r.URL.Path, "/admin/events"):
		target = eventsServiceURL
	case strings.HasPrefix(r.URL.Path, "/events"):
		target = eventsServiceURL
	default:
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	url, err := url.Parse(target)
	if err != nil {
		http.Error(w, "Bad service URL", http.StatusInternalServerError)
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(url)
	r.Host = url.Host
	proxy.ServeHTTP(w, r)
}
