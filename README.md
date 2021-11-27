# Golib GIN

Gin Wrapper to adapt with [GoLib](https://gitlab.com/golibs-starter/golib)

### Setup instruction

Base setup, see [GoLib Instruction](https://gitlab.com/golibs-starter/golib/-/blob/develop/README.md)

Both `go get` and `go mod` are supported.

```shell
go get gitlab.com/golibs-starter/golib-gin
```

### Usage

Using `fx.Option` to include dependencies for injection.

```go
package main

import (
	"gitlab.com/golibs-starter/golib-gin"
	"go.uber.org/fx"
)

func main() {
	options := []fx.Option{
		// When you want to attach Golib context to your Gin application
		golibgin.RegisterContextOpt(),

		// When you want to use default starting invoker
		golibgin.StartOpt(),

		// When you want to use default starting invoker in test
		golibgin.StartTestOpt(func(err error) {
			// Handle done state
		}),
	}
}
```