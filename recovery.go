package golibgin

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gitlab.com/golibs-starter/golib/web/log"
	"net"
	"net/http"
	"os"
	"strings"
)

func Recovery() gin.HandlerFunc {
	return CustomRecovery(DefaultHandleRecovery)
}

// CustomRecovery returns a middleware that recovers from any panics and calls the provided handle func to handle it.
func CustomRecovery(handle gin.RecoveryFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					var se *os.SyscallError
					if errors.As(ne, &se) {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}
				log.Error(c.Request.Context(), "[Recovery] Panic recovered: %s", err)
				if brokenPipe {
					// If the connection is dead, we can't write a status to it.
					_ = c.Error(err.(error))
					c.Abort()
				} else {
					handle(c, err)
				}
			}
		}()
		c.Next()
	}
}

func DefaultHandleRecovery(c *gin.Context, err interface{}) {
	c.AbortWithStatus(http.StatusInternalServerError)
}
