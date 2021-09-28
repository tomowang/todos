package middleware

import (
	"time"

	"todos/util"

	"github.com/gin-gonic/gin"
	"github.com/phuslu/log"
)

// JSONLogMiddleware logs a gin HTTP request in JSON format, with some additional custom key/values
func JSONLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()

		// Process Request
		c.Next()

		// Stop timer
		duration := util.GetDurationInMillseconds(start)

		var logger *log.Entry
		if c.Writer.Status() >= 500 {
			logger = log.Error()
		} else {
			logger = log.Info()
		}
		logger.
			Str("client_ip", util.GetClientIP(c)).
			Float64("duration", duration).
			Str("method", c.Request.Method).
			Str("path", c.Request.RequestURI).
			Int("status", c.Writer.Status()).
			Str("referrer", c.Request.Referer()).
			Str("orgion", c.Request.Header.Get("Origin")).
			Str("request_id", c.Writer.Header().Get("X-Request-Id"))
		logger.Msg("gin")
	}
}
