package main

import (
	"regexp"
	"testing"

	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_mainHandler(t *testing.T) {
	type want struct {
		statusCode int
	}
	tests := []struct {
		name    string
		method  string
		request string
		body    string
		want    want
	}{
		{
			name:    "POST request without body",
			request: "/",
			method:  http.MethodPost,
			body:    "",
			want: want{
				statusCode: http.StatusBadRequest,
			},
		},

		{
			name:    "POST request wwith body",
			request: "/",
			method:  http.MethodPost,
			body:    "https://yandex.ru",
			want: want{
				statusCode: http.StatusCreated,
			},
		},

		{
			name:    "GET request when only POST allowed",
			request: "/",
			method:  http.MethodGet,
			want: want{
				statusCode: http.StatusBadRequest,
			},
		},

		{
			name:    "POST request when only GET allowed",
			request: "/1karp",
			method:  http.MethodPost,
			want: want{
				statusCode: http.StatusBadRequest,
			},
		},

		{
			name:    "GET request incorrect id ",
			request: "/1karp",
			method:  http.MethodGet,
			want: want{
				statusCode: http.StatusBadRequest,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			body := strings.NewReader(tt.body)
			request := httptest.NewRequest(tt.method, tt.request, body)
			w := httptest.NewRecorder()
			mainHandler(w, request)
			result := w.Result()

			err := result.Body.Close()
			require.NoError(t, err)

			assert.Equal(t, result.StatusCode, tt.want.statusCode)

		})
	}
}

func Test_randSeqGen(t *testing.T) {
	tests := []struct {
		name string
		want *regexp.Regexp
	}{
		{
			name: "valid short URL",
			want: regexp.MustCompile(`^.{8}$`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Regexp(t, tt.want, randSeqGen())
		})
	}
}
