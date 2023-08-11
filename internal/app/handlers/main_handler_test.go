package handlers

import (
	"strings"
	"testing"

	"net/http"
	"net/http/httptest"

	"github.com/1karp/link_shortener/internal/app/storage"
	"github.com/stretchr/testify/require"
)

func Test_MainHandler(t *testing.T) {
	tests := []struct {
		name    string
		method  string
		request string
		body    string
		want    int
	}{
		{
			name:    "POST request without body",
			request: "/",
			method:  http.MethodPost,
			body:    "",
			want:    http.StatusBadRequest,
		},
		{
			name:    "POST request with body",
			request: "/",
			method:  http.MethodPost,
			body:    "https://yandex.ru",
			want:    http.StatusCreated,
		},
		{
			name:    "GET request when only POST allowed",
			request: "/",
			method:  http.MethodGet,
			want:    http.StatusBadRequest,
		},
		{
			name:    "POST request when only GET allowed",
			request: "/1karp",
			method:  http.MethodPost,
			want:    http.StatusBadRequest,
		},
		{
			name:    "GET request with incorrect id",
			request: "/1karp",
			method:  http.MethodGet,
			want:    http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body := strings.NewReader(tt.body)
			request := httptest.NewRequest(tt.method, tt.request, body)
			w := httptest.NewRecorder()
			MainHandler(w, request, storage.NewStorage())
			result := w.Result()

			err := result.Body.Close()
			require.NoError(t, err)

			require.Equal(t, tt.want, result.StatusCode)
		})
	}
}
