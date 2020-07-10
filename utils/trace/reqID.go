// Copyright © 2020 Bin Liu <bin.liu@enmotech.com>

package trace

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// ContextKey is
type ContextKey string

const (
	// ContextKeyReqID is the context key for RequestID
	ContextKeyReqID ContextKey = "requestID"

	// HTTPHeaderNameRequestID has the name of the header for request ID
	HTTPHeaderNameRequestID = "X-Request-ID"
)

// GetReqID will get reqID from a http request and return it as a string
func GetReqID(ctx context.Context) string {

	reqID := ctx.Value(ContextKeyReqID)

	if ret, ok := reqID.(string); ok {
		return ret
	}

	return ""
}

// AttachReqIDWithReqID will attach a brand new request ID to a http request
func AttachReqIDWithReqID(ctx context.Context, reqID string) context.Context {
	return context.WithValue(ctx, ContextKeyReqID, reqID)
}

// AttachReqID will attach a brand new request ID to a http request
func AttachReqID(ctx context.Context) context.Context {

	reqID := uuid.New()

	return context.WithValue(ctx, ContextKeyReqID, reqID.String())
}

// Middleware will attach the reqID to the http.Request and add reqID to http header in the response
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ctx := AttachReqID(r.Context())

		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)

		h := w.Header()

		h.Add(HTTPHeaderNameRequestID, GetReqID(ctx))
	})
}

// RequestID is a middleware that injects a request ID into the context of each
// request. A request ID is a string of the form "host.example.com/random-0001",
// where "random" is a base62 random string that uniquely identifies this go
// process, and where the last number is an atomically incremented request
// counter.
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := AttachReqID(c.Request.Context())
		r := c.Request.WithContext(ctx)
		c.Request = r
		c.Next()
	}
}
