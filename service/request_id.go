package service

import (
	"context"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

type contextKey string

var requestIDContextKey contextKey = "requestID"

func WithGinContext(parent *gin.Context) context.Context {
	return context.WithValue(parent, requestIDContextKey, requestid.Get(parent))
}

func WithRequestID(parent context.Context, requestID string) context.Context {
	return context.WithValue(parent, requestIDContextKey, requestID)
}

func GetRequestID(ctx context.Context) string {
	requestID, ok := ctx.Value(requestIDContextKey).(string)
	if !ok {
		return ""
	}
	return requestID
}
