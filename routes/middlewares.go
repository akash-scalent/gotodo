package routes

import (
	"context"
	"net/http"

	"github.com/akash-scalent/gotodo/utils"
	"github.com/rs/xid"
)

// AttachRequestID will attach a brand new request ID to a http request
func assignRequestID(ctx context.Context) context.Context {

	reqID := xid.New()

	return context.WithValue(ctx, utils.ContextKeyRequestID, reqID.String())
}

func reqIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		r = r.WithContext(assignRequestID(ctx))
		routerLogger.Info().Str("method", r.Method).Str("requestURI", r.RequestURI).Str("remoteAddr", r.RemoteAddr).Str("requestID",utils.GetRequestID(r.Context())).Msg("Incoming request")

		next.ServeHTTP(w, r)

		routerLogger.Info().Str("requestID", utils.GetRequestID(r.Context())).Msg("Finished http req")
	})
}
