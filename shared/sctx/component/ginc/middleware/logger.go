package middleware

import (
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"go.opentelemetry.io/otel/trace"
)

func Logger() gin.HandlerFunc {
	return logger.SetLogger(
		logger.WithLogger(func(c *gin.Context, l zerolog.Logger) zerolog.Logger {
			if trace.SpanFromContext(c.Request.Context()).SpanContext().IsValid() {
				l = l.With().
					Str("trace_id", trace.SpanFromContext(c.Request.Context()).SpanContext().TraceID().String()).
					Str("span_id", trace.SpanFromContext(c.Request.Context()).SpanContext().SpanID().String()).
					Logger()
			}

			return l.Output(gin.DefaultWriter).With().Logger()
		}),
	)
}
