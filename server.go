package golibgin

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"gitlab.com/golibs-starter/golib"
	"gitlab.com/golibs-starter/golib/log"
	"go.uber.org/fx"
	"net/http"
)

func GinHttpServerOpt() fx.Option {
	return fx.Options(
		fx.Provide(NewGinEngine),
		fx.Provide(NewHTTPServer),
		fx.Invoke(RegisterHandlers),
		fx.Invoke(OnStartHttpServerHook),
	)
}

func OnStopHttpServerOpt() fx.Option {
	return fx.Invoke(OnStopHttpServerHook)
}

type GinEngineIn struct {
	fx.In
	Logging *log.Properties `optional:"true"`
}

func NewGinEngine(in GinEngineIn) *gin.Engine {
	if in.Logging != nil && in.Logging.Development {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	engine := gin.New()
	engine.Use(Recovery())
	return engine
}

func NewHTTPServer(app *golib.App, engine *gin.Engine) *http.Server {
	return &http.Server{
		Addr:    fmt.Sprintf(":%d", app.Port()),
		Handler: engine,
	}
}

func RegisterHandlers(app *golib.App, engine *gin.Engine) {
	engine.Use(InitContext())
	engine.Use(WrapAll(app.Handlers())...)
}

func OnStartHttpServerHook(lc fx.Lifecycle, app *golib.App, httpServer *http.Server) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Infof("Application will be served at %s. Service name: %s, service path: %s",
				httpServer.Addr, app.Name(), app.Path())
			go func() {
				if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					log.Errorf("Could not serve HTTP request at %s, error [%v]", httpServer.Addr, err)
				}
				log.Infof("Stopped HTTP Server %s", httpServer.Addr)
			}()
			return nil
		},
	})
}

func OnStopHttpServerHook(lc fx.Lifecycle, httpServer *http.Server) {
	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			log.Infof("Stopping HTTP Server %s", httpServer.Addr)
			if err := httpServer.Shutdown(ctx); err != nil {
				log.Errorf("Could not stop HTTP server, error [%v]", err)
			}
			return nil
		},
	})
}

// StartOpt
// Deprecated: Using GinHttpServerOpt in bootstrap options instead
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
