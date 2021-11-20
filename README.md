# Golib GIN

Gin Wrapper to adapt with [GoLib](https://gitlab.id.vin/vincart/golib)

### Setup instruction

Base setup, see [GoLib Instruction](https://gitlab.id.vin/vincart/golib/-/blob/develop/README.md)

Both `go get` and `go mod` are supported.

```shell
go get gitlab.id.vin/vincart/golib-gin
```

### Usage

Using `fx.Option` to include dependencies for injection.

```go
package main

import (
	"gitlab.id.vin/vincart/golib-gin"
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