package codes

import (
	"fmt"
	"log"
	"net/http"
)

type StatCode int

const (
	SERVER_OK        StatCode = 200
	SERVER_ERR       StatCode = 500
	SERVER_BAD_INPUT StatCode = 400
)

// Report the message: log it, write the Status code header, write the message.
func (s StatCode) Report(w http.ResponseWriter, msg string) {
	log.Println(msg)
	w.WriteHeader(int(s))
	_, _ = fmt.Fprintln(w, msg)
}

// Check for an error, and if any: log the message, log the error, write the status code, write the message.
// Returns true if there was an error, false otherwise.
//
// Error handling is now just a simple if MY_CODE.Check(w, err, "foobar errored") { return }
func (s StatCode) Check(w http.ResponseWriter, err error, msg string) bool {
	if err != nil {
		log.Println(msg)
		log.Println(err)
		w.WriteHeader(int(s))
		_, _ = fmt.Fprintln(w, msg)
		return true
	} else {
		return false
	}
}
