// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package pprof

import (
	"net/http/pprof"
	"strings"

	"clevergo.tech/clevergo"
)

var (
	index    = clevergo.HandleHandlerFunc(pprof.Index)
	cmdline  = clevergo.HandleHandlerFunc(pprof.Cmdline)
	profile  = clevergo.HandleHandlerFunc(pprof.Profile)
	symbol   = clevergo.HandleHandlerFunc(pprof.Symbol)
	trace    = clevergo.HandleHandlerFunc(pprof.Trace)
	handlers = map[string]clevergo.Handle{
		"cmdline": cmdline,
		"profile": profile,
		"symbol":  symbol,
		"trace":   trace,
	}
)

// RegisterHandler register pprof handler.
func RegisterHandler(router clevergo.Router) {
	router.Any("/debug/pprof/*name", Handler)
}

// Handler wraps all pprof handlers.
func Handler(c *clevergo.Context) error {
	name := strings.TrimPrefix(c.Params.String("name"), "/")
	if handler, ok := handlers[name]; ok {
		return handler(c)
	}
	return index(c)
}
