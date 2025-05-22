package sentrymiddleware

import (
	"context"
	"fmt"
	"github.com/getsentry/sentry-go"
	sentryhttp "github.com/getsentry/sentry-go/http"
	"github.com/rs/zerolog"
	"net/http"
)

func Sentry(dsn string) (func(http.Handler) http.Handler, error) {
	sentryOptions := sentry.ClientOptions{
		Dsn: dsn,
		BeforeSend: func(event *sentry.Event, hint *sentry.EventHint) *sentry.Event {
			if req, ok := hint.Context.Value(sentry.RequestContextKey).(*http.Request); ok {
				// Add IP Address to user information.
				event.User.IPAddress = req.RemoteAddr
			}

			return event
		},
	}

	if err := sentry.Init(sentryOptions); err != nil {
		return nil, fmt.Errorf("sentry.Init: %w", err)
	}

	sentryHandler := sentryhttp.New(sentryhttp.Options{})

	return sentryHandler.Handle, nil
}

func CaptureError(ctx context.Context, err error) {
	logger := zerolog.Ctx(ctx)
	hub := sentry.GetHubFromContext(ctx)

	logger.Error().Err(err).Msg("")

	if hub != nil {
		hub.CaptureException(err)
	}
}

func CaptureMessage(ctx context.Context, message string) {
	logger := zerolog.Ctx(ctx)
	hub := sentry.GetHubFromContext(ctx)

	logger.Error().Msg(message)

	if hub != nil {
		hub.CaptureMessage(message)
	}
}
