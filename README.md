# Golib GIN

> **Note**
> We are moving out from [Gitlab](https://gitlab.com/golibs-starter). All packages are now migrated
> to `github.com/golibs-starter/*`. Please consider updating.

Gin Wrapper to adapt with [GoLib](https://github.com/golibs-starter/golib)

### Setup instruction

Base setup, see [GoLib Instruction](https://github.com/golibs-starter/golib#readme)

Both `go get` and `go mod` are supported.

```shell
go get github.com/golibs-starter/golib-gin
```

### Usage

Using `fx.Option` to include dependencies for injection.

```go
package main

import (
    "github.com/golibs-starter/golib-gin"
    "go.uber.org/fx"
)

func main() {
    fx.New(
        // Using Gin as handler for Http Server,
        // Append startup method to Fx OnStart lifecycle.
        golibgin.GinHttpServerOpt(),

        // When you want to enable graceful shutdown.
        golibgin.OnStopHttpServerOpt(),
    ).Run()
}
```
