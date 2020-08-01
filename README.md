# CleverGo pprof
[![Build Status](https://img.shields.io/travis/clevergo/pprof?style=for-the-badge)](https://travis-ci.org/clevergo/pprof)
[![Coverage Status](https://img.shields.io/coveralls/github/clevergo/pprof?style=for-the-badge)](https://coveralls.io/github/clevergo/pprof)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white&style=for-the-badge)](https://pkg.go.dev/clevergo.tech/pprof?tab=doc)
[![Go Report Card](https://goreportcard.com/badge/github.com/clevergo/pprof?style=for-the-badge)](https://goreportcard.com/report/github.com/clevergo/pprof)
[![Release](https://img.shields.io/github/release/clevergo/pprof.svg?style=for-the-badge)](https://github.com/clevergo/pprof/releases)
[![Downloads](https://img.shields.io/endpoint?url=https://pkg.clevergo.tech/api/badges/downloads/month/clevergo.tech/pprof&style=for-the-badge)](https://pkg.clevergo.tech/)

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
    pprof.RegisterHandlers(app)
    app.Run(":8080")
}
```
