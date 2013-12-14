// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package handlerutil

import (
	"net/http"
)

func GetRequestedPathFromRequest(r *http.Request) string {
	requestedPath := r.URL.Path
	return requestedPath
}

func GetHostnameFromRequest(r *http.Request) string {
	return r.Host
}