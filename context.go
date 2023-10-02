package golibgin

import (
	"github.com/gin-gonic/gin"
	"github.com/golibs-starter/golib/web/constant"
	"github.com/golibs-starter/golib/web/context"
)

func InitContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestAttributes := context.GetOrCreateRequestAttributes(c.Request)
		requestAttributes.Mapping = c.FullPath()
		c.Set(constant.ContextReqAttribute, requestAttributes)
		c.Next()
	}
}
