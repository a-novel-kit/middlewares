package golmmiddleware

import (
	"net/http"

	"github.com/a-novel-kit/golm"
)

func Golm[Req, Res, Stream any](binding golm.ChatBinding[Req, Res, Stream]) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			// Inject the chat into the following middlewares context.
			next.ServeHTTP(w, r.WithContext(golm.WithContext(ctx, binding)))
		})
	}
}
