package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/pkg/errors"
)

// ---
// main and run
// https://youtu.be/rWBSMsLG8po?t=552

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() (err error) {
	db, dbtidy, err := setupDatabase()
	if err != nil {
		return errors.Wrap(err, "setup database")
	}
	defer dbtidy()
	srv := &server{
		db:     db,
		router: &router{},
		email:  emailSender{},
	}
	// ... more stuff

	// ---
	// ANCILLARY (to make compiler happy)
	srv.ServeHTTP(&responseWriter{}, &http.Request{})
	return nil
}

// ---
// ANCILLARY STUBS (to make compiler happy)

func setupDatabase() (db *dbConn, dbtidy func(), err error) {
	db = &dbConn{}
	dbtidy = func() {}
	return db, dbtidy, err
}

// responseWriter is an implementation of http.ResponseWriter
type responseWriter struct{}

func (w *responseWriter) Header() (head http.Header)         { return head }
func (w *responseWriter) Write([]byte) (code int, err error) { return code, err }
func (w *responseWriter) WriteHeader(statusCode int)         {}
