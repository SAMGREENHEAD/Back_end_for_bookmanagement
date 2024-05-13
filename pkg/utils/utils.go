// To marshall and unmarshall data

package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// ParseBody reads the request body of an HTTP request and unmarshals it to the provided interface{}
// The function is useful for unmarshalling JSON data from the request body into a Go struct.
//
// Parameters:
//  r: The HTTP request object
//  x: The interface{} to unmarshal the request body into
//
// Returns: None
func ParseBody(r *http.Request, x interface{}) {
	// Read the request body
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		// Attempt to unmarshal the request body into the provided interface
		if err := json.Unmarshal(body, x); err != nil {
			// Return immediately if there is an error
			return
		}
	}
}
