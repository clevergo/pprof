// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package pprof

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"clevergo.tech/clevergo"
	"github.com/stretchr/testify/assert"
)

func TestRegisterHandler(t *testing.T) {
	app := clevergo.Pure()
	RegisterHandler(app)

	// index
	req := httptest.NewRequest(http.MethodGet, "/debug/pprof/", nil)
	resp := httptest.NewRecorder()
	app.ServeHTTP(resp, req)
	assert.Contains(t, resp.Body.String(), "<title>/debug/pprof/</title>")

	// profile
	req = httptest.NewRequest(http.MethodGet, "/debug/pprof/profile", nil)
	resp = httptest.NewRecorder()
	app.ServeHTTP(resp, req)
	assert.Equal(t, resp.Header().Get("Content-Disposition"), `attachment; filename="profile"`)
}
