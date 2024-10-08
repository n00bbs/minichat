package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/nineteenseventy/minichat/core/logging"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func LoggerMiddleware() func(http.Handler) http.Handler {
	logger := logging.GetLogger("http.middleware.logger")

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			t1 := time.Now()

			lrw := loggingResponseWriter{
				ResponseWriter: w,
				statusCode:     http.StatusOK,
			}

			defer func() {
				scheme := "http"
				if r.TLS != nil {
					scheme = "https"
				}
				logger.Debug().
					Str("from", r.RemoteAddr).
					Int("status", lrw.statusCode).
					Str("size", w.Header().Get("Content-Length")).
					Str("method", r.Method).
					Str("proto", r.Proto).
					TimeDiff("time", time.Now(), t1).
					Msg(fmt.Sprintf(
						"%s://%s%s",
						scheme,
						r.Host,
						r.RequestURI,
					))
			}()

			next.ServeHTTP(&lrw, r)
		})
	}
}
