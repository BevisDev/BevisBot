package middleware

import (
	"bytes"
	"fmt"
	"github.com/BevisDev/godev/consts"
	"github.com/BevisDev/godev/utils"
	"github.com/gin-gonic/gin"
	"io"
	"time"
)

type ResponseWrapper struct {
	gin.ResponseWriter
	body     *bytes.Buffer
	status   int
	duration float64
}

func (w *ResponseWrapper) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func loggerHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			startTime   = time.Now()
			state       = utils.GetState(c.Request.Context())
			contentType = c.Request.Header.Get(consts.ContentType)
		)

		// ignore log some content-type
		skipBody := utils.SkipContentType(contentType)

		// log request
		var reqBody string
		if !skipBody {
			reqBytes, _ := io.ReadAll(c.Request.Body)
			reqBody = string(reqBytes)
			c.Request.Body = io.NopCloser(bytes.NewBuffer(reqBytes))
		}

		fmt.Println(state, reqBody)
		
		//lib.Logger.LogRequest(&logger.RequestLogger{
		//	State:       state,
		//	URL:         c.Request.URL.String(),
		//	RequestTime: startTime,
		//	Query:       c.Request.URL.RawQuery,
		//	Method:      c.Request.Method,
		//	Body:        reqBody,
		//})

		// wrap the responseWriter to capture the response body
		respBuffer := &bytes.Buffer{}
		writer := &ResponseWrapper{
			ResponseWriter: c.Writer,
			body:           respBuffer,
		}
		c.Writer = writer

		// process next
		c.Next()

		// log response
		var (
			duration = time.Since(startTime)
			respBody string
		)
		contentType = c.Writer.Header().Get(consts.ContentType)
		skipBody = utils.SkipContentType(contentType)
		if !skipBody {
			respBody = writer.body.String()
		}

		fmt.Println(duration, respBody)

		//lib.Logger.LogResponse(&logger.ResponseLogger{
		//	State:       state,
		//	Status:      c.Writer.Status(),
		//	DurationSec: duration,
		//	Body:        respBody,
		//})
	}
}
