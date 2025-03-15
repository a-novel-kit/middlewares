package golmmiddleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/a-novel-kit/golm"
	golmmocks "github.com/a-novel-kit/golm/mocks"

	golmmiddleware "github.com/a-novel-kit/middlewares/golm"
)

func TestGolm(t *testing.T) {
	t.Parallel()

	binding := golmmocks.NewMockChatBinding[string, string, string](t)
	binding.EXPECT().
		Completion(
			mock.Anything,
			mock.AnythingOfType("golm.UserMessage"),
			mock.AnythingOfType("golm.CompletionParams"),
			mock.AnythingOfType("golm.ChatHistory"),
		).
		Return(&golm.AssistantMessage{Content: "test"}, nil)

	baseHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		chat := golm.Context(r.Context())

		_, err := chat.Completion(r.Context(), golm.NewUserMessage("foo"), golm.CompletionParams{})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			return
		}

		w.WriteHeader(http.StatusOK)
	})
	baseHandlerWithChat := golmmiddleware.Golm[string, string, string](binding)(baseHandler)

	req := httptest.NewRequest(http.MethodGet, "/something", nil)

	w := httptest.NewRecorder()
	baseHandlerWithChat.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
}
