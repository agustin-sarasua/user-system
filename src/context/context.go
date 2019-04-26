package context

import (
	"context"

	"github.com/gin-gonic/gin"
)

type ctxKey int

const (
	requestCtx   = "request_context"
	uuidKey      = ctxKey(iota)
	TokenAuthKey = "X-Auth-Token"
)

//WithRequestContext returns a context.Context from a gin.Context
func WithRequestContext(ctx context.Context, ginContext *gin.Context) {
	ginContext.Set(requestCtx, ctx)
}

//RequestContext returns a context.Context from a gin.Context
func RequestContext(c *gin.Context) context.Context {
	ctxValue, ok := c.Get(requestCtx)
	if !ok {
		return context.Background()
	}

	return ctxValue.(context.Context)
}

//WithUUID adds a UUID for the request
func WithUUID(c context.Context, id string) context.Context {

	return context.WithValue(c, uuidKey, id)
}

//WithTokenAuth adds the Auth-Token for the request
func WithTokenAuth(c context.Context, token string) context.Context {
	return context.WithValue(c, TokenAuthKey, token)
}

//UUID gets the request's UUID
func UUID(ctx context.Context) string {
	value := ctx.Value(uuidKey)
	uuid, ok := value.(string)

	if !ok {
		return "not_defined"
	}

	return uuid
}
