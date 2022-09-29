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
    fx.New(
        // Using Gin as handler for Http Server,
        // Append startup method to Fx OnStart lifecycle.
        golibgin.GinHttpServerOpt(),
    ).Run()
}
```
