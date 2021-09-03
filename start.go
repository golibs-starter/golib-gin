package golibgin

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"gitlab.id.vin/vincart/golib"
	"gitlab.id.vin/vincart/golib/log"
	"go.uber.org/fx"
)

func StartOpt() fx.Option {
	return fx.Invoke(func(lc fx.Lifecycle, app *golib.App, engine *gin.Engine) {
		lc.Append(fx.Hook{
			OnStart: func(ctx context.Context) error {
				log.Infof("Application will be served at %d. Service name: %s, service path: %s",
					app.Port(), app.Name(), app.Path())
				go func() {
					if err := engine.Run(fmt.Sprintf(":%d", app.Port())); err != nil {
						log.Fatalf("Cannot start application due by error [%v]", err)
					}
				}()
				return nil
			},
		})
	})
}
