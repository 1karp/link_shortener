package logging

import (
	"net/http"
	"time"

	"go.uber.org/zap"
)

var Sugar *zap.SugaredLogger

func CustomMiddlewareLogger(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		h.ServeHTTP(w, r)

		Sugar.Infoln("URI:", r.RequestURI,
			"Method:", r.Method,
			"Duration:", time.Since(start))
	}

	return http.HandlerFunc(fn)
}
