package utils

import (
	"context"
)
type contextKey string

// ContextKeyRequestID is the ContextKey for RequestID
const ContextKeyRequestID contextKey = "requestID"
// GetRequestID will get reqID from a http request and return it as a string
func GetRequestID(ctx context.Context) string {
	reqID := ctx.Value(ContextKeyRequestID)
	if ret, ok := reqID.(string); ok {
		return ret
	}

	return ""
}