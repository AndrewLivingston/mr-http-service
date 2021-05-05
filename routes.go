package main

import "net/http"

// ---
// routes file

// routes in routes.go is a convenient one-stop shop for urls during
// maintenance
func (s *server) routes() {
	s.router.Get("/api/", s.handleAPI())
	s.router.Get("/about", s.handleAbout())
	s.router.Get("/", s.handleIndex())
}

//---
// ANCILLARY STUBS (to make compiler happy)

// router is an arbitrary router stub
type router struct{}

func (rr *router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// actual implementation for server.ServeHTTP
}

func (rr *router) Get(endpoint string, handler http.HandlerFunc) {}

// unexplained; used in https://youtu.be/rWBSMsLG8po?t=1080:
func (rr *router) HandleFunc(endpoint string, h http.HandlerFunc) {}

func (s *server) handleAPI() (h http.HandlerFunc)   { return h }
func (s *server) handleAbout() (h http.HandlerFunc) { return h }
func (s *server) handleIndex() (h http.HandlerFunc) { return h }
