# errors [![GoDoc](https://godoc.org/github.com/peczenyj/errors?status.svg)](http://godoc.org/github.com/peczenyj/errors) [![Report card](https://goreportcard.com/badge/github.com/peczenyj/errors)](https://goreportcard.com/report/github.com/peczenyj/errors) [![Go](https://github.com/peczenyj/errors/actions/workflows/go.yml/badge.svg)](https://github.com/peczenyj/errors/actions/workflows/go.yml)

nice replacement for [pkg/errors](https://github.com/pkg/errors)

Fork of [xerrors](https://pkg.go.dev/golang.org/x/xerrors) with explicit [Wrap](https://pkg.go.dev/github.com/peczenyj/errors#Wrap) instead of `%w` with no stack trace.

```bash
go get github.com/peczenyj/errors
```

```go
errors.Wrap(err, "message")
```

## Why

* Using `Wrap` is the most explicit way to wrap errors
* Wrapping with `fmt.Errorf("foo: %w", err)` is implicit, redundant and error-prone
* Parsing `"foo: %w"` is implicit, redundant and slow
* The [pkg/errors](https://github.com/pkg/errors) and [xerrors](https://pkg.go.dev/golang.org/x/xerrors) are not maintainted
* The [cockroachdb/errors](https://github.com/cockroachdb/errors) is too big
* The `errors` has no `Wrap`
* The [go-faster/errors](https://github.com/go-faster/errors) demands call `errors.DisableTrace` or build tag `noerrtrace` to disable stack traces.

