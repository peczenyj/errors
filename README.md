# errors 
[![GoDoc](https://godoc.org/github.com/peczenyj/errors?status.svg)](http://godoc.org/github.com/peczenyj/errors)
[![Go](https://github.com/peczenyj/errors/actions/workflows/go.yml/badge.svg)](https://github.com/peczenyj/errors/actions/workflows/go.yml)
[![Lint](https://github.com/peczenyj/errors/actions/workflows/lint.yml/badge.svg)](https://github.com/peczenyj/errors/actions/workflows/lint.yml)
[![CodeQL](https://github.com/peczenyj/errors/actions/workflows/github-code-scanning/codeql/badge.svg)](https://github.com/peczenyj/errors/actions/workflows/github-code-scanning/codeql)
[![Dependency Review](https://github.com/peczenyj/errors/actions/workflows/dependency-review.yml/badge.svg)](https://github.com/peczenyj/errors/actions/workflows/dependency-review.yml)
[![Report card](https://goreportcard.com/badge/github.com/peczenyj/errors)](https://goreportcard.com/report/github.com/peczenyj/errors)
[![codecov](https://codecov.io/gh/peczenyj/errors/graph/badge.svg?token=9y6f3vGgpr)](https://codecov.io/gh/peczenyj/errors)

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

If you don't need stack traces, this is the right tool for you.

The motivation behind this package is: if you are using [pkg/errors](https://github.com/pkg/errors) or [go-faster/errors](https://github.com/go-faster/errors) should be easier to swich to this package by just change the import (when you can't switch to the standard `errors`).

It means, if your code is already using method such as `Wrap` or `Wrapf`, they are available.

Same for `WithMessage` and `WithMessagef` (however, here it is just an alias to `Wrap` and `Wrapf`).

The useful function `Into` from `go-faster/errors` is also available.

### werrors

You can keep using standard `errors` and import the `github.com/peczenyj/errors/werrors` to have access to some extra functions such as `Wrap`, `Wrapf`, `WithMessage`, `WithMessagef`, `Cause` and `Into`.

It is a more minimalistic approach.

```bash
go get github.com/peczenyj/errors/werrors
```

```go
werrors.Wrap(err, "message")
```