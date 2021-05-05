package main

import "net/http"

// ---
// server type
// https://youtu.be/rWBSMsLG8po?t=604

type server struct {
	// eliminates global variables
	db     *dbConn
	router *router
	email  EmailSender
}

// newServer is inconsistent with example in main.go. Matt prefers not to have
// constructors, but says he always seems to end up with one for servers
func newServer() *server {
	s := &server{} // leaves dependencies as nil (could be passed as
	// arguments if not many)
	s.routes()
	return s
}

// ---
// server as http.Handler
// https://youtu.be/rWBSMsLG8po?t=772

// ServeHTTP makes server an http.Handler. (Recall this method is intended to
// panic instead of returning an error.)
func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r) // delegates to router for actual implementation
	// don't put logic in here, use middleware instead
}

// ---
// multiple server types
// https://youtu.be/rWBSMsLG8po?t=1171

// a pattern for large apps; can have different dependencies

// in people.go
type serverPeople struct {
	db    *dbConn
	email EmailSender
}

// comments.go
type serverComments struct {
	db *dbConn
	// lacks EmailSender dependencies
}

// ---
// ANCILLARY STUBS (to make compiler happy)

// Conn is a an arbitrary database connection stub
type dbConn struct{}

// EmailSender is probably an interface because name is verb
type EmailSender interface {
	Send() error
}

// emailSender is an implementation of EmailSender
type emailSender struct{}

// Send is a stub function to give Sender something to match EmailSender
func (e emailSender) Send() (err error) { return err }
