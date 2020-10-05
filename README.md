# CleverGo pprof
[![Build Status](https://img.shields.io/travis/clevergo/pprof?style=flat-square)](https://travis-ci.org/clevergo/pprof)
[![Coverage Status](https://img.shields.io/coveralls/github/clevergo/pprof?style=flat-square)](https://coveralls.io/github/clevergo/pprof)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/clevergo.tech/pprof?tab=doc)
[![Go Report Card](https://goreportcard.com/badge/github.com/clevergo/pprof?style=flat-square)](https://goreportcard.com/report/github.com/clevergo/pprof)
[![Release](https://img.shields.io/github/release/clevergo/pprof.svg?style=flat-square)](https://github.com/clevergo/pprof/releases)
[![Downloads](https://img.shields.io/endpoint?url=https://pkg.clevergo.tech/api/badges/downloads/total/clevergo.tech/pprof&style=flat-square)](https://pkg.clevergo.tech/)
[![Chat](https://img.shields.io/badge/chat-telegram-blue?style=flat-square)](https://t.me/clevergotech)
[![Community](https://img.shields.io/badge/community-forum-blue?style=flat-square&color=orange)](https://forum.clevergo.tech)

```shell
$ go get -u clevergo.tech/pprof
```

```go
import (
    "clevergo.tech/clevergo"
    "clevergo.tech/pprof"
)

func main() {
    app := clevergo.New()
    pprof.RegisterHandler(app)
    app.Run(":8080")
}
```
