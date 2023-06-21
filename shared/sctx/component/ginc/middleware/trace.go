package middleware

import (
	"bytes"
	"encoding/json"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/trace"
)

func Traceable() gin.HandlerFunc {
	return func(c *gin.Context) {
		wb := &BodyWriter{
			body:           &bytes.Buffer{},
			ResponseWriter: c.Writer,
		}
		c.Writer = wb

		c.Next()

		if trace.SpanFromContext(c.Request.Context()).SpanContext().IsValid() {
			traceId := trace.SpanFromContext(c.Request.Context()).SpanContext().TraceID().String()

			res := make(map[string]interface{})
			err := json.Unmarshal(wb.body.Bytes(), &res)
			if err == nil {
				res["trace_id"] = traceId
			}
			wb.body = &bytes.Buffer{}

			bytes, err := json.Marshal(res)
			if err == nil {
				wb.Write(bytes)
			}
		}

		wb.ResponseWriter.Write(wb.body.Bytes())
	}
}

type BodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r BodyWriter) Write(b []byte) (int, error) {
	return r.body.Write(b)
}
