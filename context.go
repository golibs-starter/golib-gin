package golibgin

import (
	"github.com/gin-gonic/gin"
	"gitlab.id.vin/vincart/golib"
	"gitlab.id.vin/vincart/golib/web/constant"
	"gitlab.id.vin/vincart/golib/web/context"
	"go.uber.org/fx"
)

func RegisterContextOpt() fx.Option {
	return fx.Invoke(RegisterContext)
}

func RegisterContext(app *golib.App, engine *gin.Engine) {
	engine.Use(InitContext())
	engine.Use(WrapAll(app.Handlers())...)
}

func InitContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestAttributes := context.GetOrCreateRequestAttributes(c.Request)
		requestAttributes.Mapping = c.FullPath()
		c.Set(constant.ContextReqAttribute, requestAttributes)
		c.Next()
	}
}
