package middlewares

import (
	"errors"
	"net/http"
	"time"

	"github.com/a-novel-kit/context"
)

// UseContext is a middleware used to override the default request of a group of handlers.
// Returns a chi-compatible middleware.
func UseContext(initial context.Context, timeout time.Duration) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx, cancel := context.WithTimeout(initial, timeout)
			defer cancel()

			// Make sure context is properly closed if the request takes too long.
			go func() {
				<-ctx.Done()

				if errors.Is(ctx.Err(), context.DeadlineExceeded) {
					w.WriteHeader(http.StatusGatewayTimeout)
				}
			}()

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
