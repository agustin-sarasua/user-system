package context

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

const requestIDHeader = "X-Request-Id"

func PropagateHeaders(c *gin.Context) {
	tokenAuth := c.Request.Header.Get("X-Auth-Token")
	requestCtx := RequestContext(c)
	if tokenAuth != "" {
		requestCtx = WithTokenAuth(requestCtx, tokenAuth)
		c.Writer.Header().Set("X-Auth-Token", tokenAuth)
	}
	requestID := c.Request.Header.Get(requestIDHeader)
	if requestID == "" {
		uuidT, _ := uuid.NewV4()
		requestID = uuidT.String()
	}
	requestCtx = WithUUID(requestCtx, requestID)
	c.Writer.Header().Set(requestIDHeader, requestID)
	WithRequestContext(requestCtx, c)
	c.Next()
}
