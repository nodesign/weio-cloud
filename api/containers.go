/**
 * Copyright (c) Mainflux
 *
 * Mainflux server is licensed under an Apache license, version 2.0.
 * All rights not explicitly granted in the Apache license, version 2.0 are reserved.
 * See the included LICENSE file for more details.
 */

package api

import (
	"io"
	"net/http"
)

/** == Functions == */

// createContainer function
func createContainer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	// Launch

	w.WriteHeader(http.StatusOK)
	str := `{"response": "created"}`
	io.WriteString(w, str)
}

// getDevices function
func getContainers(w http.ResponseWriter, r *http.Request) {
	println("Get Containers")
}
