# errors

[![tag](https://img.shields.io/github/tag/peczenyj/errors.svg)](https://github.com/peczenyj/errors/releases)
![Go Version](https://img.shields.io/badge/Go-%3E%3D%201.18-%23007d9c)
[![GoDoc](https://pkg.go.dev/badge/github.com/peczenyj/errors)](http://pkg.go.dev/github.com/peczenyj/errors)
[![Go](https://github.com/peczenyj/errors/actions/workflows/go.yml/badge.svg)](https://github.com/peczenyj/errors/actions/workflows/go.yml)
[![Lint](https://github.com/peczenyj/errors/actions/workflows/lint.yml/badge.svg)](https://github.com/peczenyj/errors/actions/workflows/lint.yml)
[![codecov](https://codecov.io/gh/peczenyj/errors/graph/badge.svg?token=9y6f3vGgpr)](https://codecov.io/gh/peczenyj/errors)
[![Report card](https://goreportcard.com/badge/github.com/peczenyj/errors)](https://goreportcard.com/report/github.com/peczenyj/errors)
[![CodeQL](https://github.com/peczenyj/errors/actions/workflows/github-code-scanning/codeql/badge.svg)](https://github.com/peczenyj/errors/actions/workflows/github-code-scanning/codeql)
[![Dependency Review](https://github.com/peczenyj/errors/actions/workflows/dependency-review.yml/badge.svg)](https://github.com/peczenyj/errors/actions/workflows/dependency-review.yml)
[![License](https://img.shields.io/github/license/peczenyj/errors)](./LICENSE)

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

## Features

| Feature                                         | `errors` | `pkg/errors` | `go-faster/errors` | `peczenyj/errors` |
|-------------------------------------------------|----------|--------------|--------------------|-------------------|
| error constructors (`New`, `Errorf`)            | ✔        | ✔            | ✔                  | ✔                 |
| error causes (`Cause` / `Unwrap`)               |          | ✔            | ✔                  | ✔                 |
| type safety (`Into`)                            |          |              | ✔                  | ✔                 |
| `errors.As()`, `errors.Is()`                    | ✔        | ✔            | ✔                  | ✔                 |
| support stack traces                            |          | ✔            | ✔                  | no, by desing     |
|

## Motivation

If your project wants to minimize external dependencies and does not need stack traces on every error or failure, this is a acceptable alternative.

When migrating from some `github.com/<your favorite repository>/errors` to the standard lib `errors`, we can see that our code rely on non-standard functions such as `Wrap` or `Wrapf` that demands more changes than just update the import.

However, helper function such as `Wrap` or `Wrapf` are useful, since it highlight the **error** and automatically ignores nil values.

It means we can choose between:

```go
data, err := io.ReadAll(r)
if err != nil {
        return errors.Wrap(err, "read failed") // add context
}
...
```

And a more simple approach:

```go
data, err := io.ReadAll(r)

return data, errors.Wrap(err, "read failed") // will return nil if err is nil
```

### werrors

You can keep using standard `errors` and import the `github.com/peczenyj/errors/werrors` to have access to some extra functions such as `Wrap`, `Wrapf`, `WithMessage`, `WithMessagef`, `Cause` and `Into`.

It is a more minimalistic approach.

```bash
go get github.com/peczenyj/errors/werrors
```

```go
werrors.Wrap(err, "message")
```
