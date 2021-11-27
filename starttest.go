package golibgin

import (
	"context"
	"github.com/gin-gonic/gin"
	"gitlab.com/golibs-starter/golib"
	"gitlab.com/golibs-starter/golib/log"
	"go.uber.org/fx"
	"net"
	"net/http"
	"time"
)

func StartTestOpt(done func(err error)) fx.Option {
	return fx.Invoke(func(lifecycle fx.Lifecycle, app *golib.App, engine *gin.Engine, ln net.Listener) {
		lifecycle.Append(fx.Hook{
			OnStart: func(ctx context.Context) error {
				go func() {
					if err := http.Serve(ln, engine); err != nil {
						log.Errorf("Cannot start application due by error [%v]", err)
					}
				}()
				time.Sleep(50 * time.Millisecond)
				done(nil)
				return nil
			},
		})
	})
}
