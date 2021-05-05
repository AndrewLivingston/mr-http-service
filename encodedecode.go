package main

import (
	"encoding/json"
	"net/http"
)

// ---
// some common (premature?) abstractions
// https://youtu.be/rWBSMsLG8po?t=1508

// always taking ResponseWriter and Request future-proofs these

// respond should start as a bare-bones abstraction and only be made more
// sophisticated as needed.
func (s *server) respond(
	w http.ResponseWriter,
	r *http.Request,
	data interface{},
	status int,
) error {
	w.WriteHeader(status)
	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			return err
		}
	}
	// ... more stuff
	return nil
}

// decode should also start as simple as possible. Mat admits this abstraction
// is unpopular
func (s *server) decode(w http.ResponseWriter, r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}
