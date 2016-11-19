/**
 * Copyright (c) Mainflux
 *
 * Mainflux server is licensed under an Apache license, version 2.0.
 * All rights not explicitly granted in the Apache license, version 2.0 are reserved.
 * See the included LICENSE file for more details.
 */

package api

import (
	"net/http"
	"strconv"

	"github.com/nodesign/weio-cloud/config"

	"github.com/go-zoo/bone"
)

// HTTPServer function
func HTTPServer(cfg config.Config) {

	mux := bone.New()

	/**
	 * Routes
	 */
	// Status
	mux.Get("/status", http.HandlerFunc(getStatus))

	// Devices
	mux.Post("/containers", http.HandlerFunc(createContainer))
	mux.Get("/containers", http.HandlerFunc(getContainers))

	/**
	 * Server
	 */
	http.ListenAndServe(cfg.HTTPHost+":"+strconv.Itoa(cfg.HTTPPort), mux)
}
