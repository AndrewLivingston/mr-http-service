package main

import "net/http"

// ---
// middleware
// https://youtu.be/rWBSMsLG8po?t=1314

// adminOnly is a decorator that simply passes along a HandlerFunc, first
// ensuring the user has admin access
func (s *server) adminOnly(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// can run code before calling wrapped handler h:
		if !currentUser(r).IsAdmin {
			http.NotFound(w, r)
			return // don't call wrapped handler h at all
		}
		h(w, r)
		// can also run code after calling wrapped handler h
	}
}

// example of using middleware on an admin route:

func (s *server) handleAdminIndex() (h http.HandlerFunc) { return h }

func (s *server) routesWithMiddleWare() {
	s.router.Get("/admin", s.adminOnly(s.handleAdminIndex()))
}

// ---
// ANCILLARY STUBS (to make compiler happy)

type user struct {
	IsAdmin bool
}

func currentUser(r *http.Request) (u user) { return u }
