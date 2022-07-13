package middlewares

import (
	"admin-manage-shop/pkg/helpers"
	"admin-manage-shop/pkg/logger"
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"io/ioutil"
	"time"
)

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		w := &responseBodyWriter{
			body:           &bytes.Buffer{},
			ResponseWriter: c.Writer,
		}
		c.Writer = w

		var requestBody []byte
		if c.Request.Body != nil {
			requestBody, _ = ioutil.ReadAll(c.Request.Body)
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestBody))
		}

		start := time.Now()
		c.Next()

		cost := time.Since(start)
		responStatus := c.Writer.Status()

		logFields := []zap.Field{
			zap.Int("status", responStatus),
			zap.String("request", c.Request.Method+" "+c.Request.URL.String()),
			zap.String("query", c.Request.URL.RawQuery),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.String("time", helpers.MicrosecondsStr(cost)),
		}

		if c.Request.Method == "POST" || c.Request.Method == "PUT" || c.Request.Method == "DELETE" {
			logFields = append(logFields, zap.String("Request Body", string(requestBody)))

			logFields = append(logFields, zap.String("Response Body", w.body.String()))
		}

		if responStatus > 400 && responStatus <= 499 {
			logger.Warn("HTTP Waring "+cast.ToString(responStatus), logFields...)
		} else if responStatus >= 500 && responStatus <= 599 {
			logger.Error("HTTP Error "+cast.ToString(responStatus), logFields...)
		} else {
			logger.Debug("HTTP Access Log", logFields...)
		}
	}
}
