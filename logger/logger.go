package logger

import (
	"context"
	"go.uber.org/zap"
	"net/http"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		l := zap.NewExample()
		l = l.With(zap.Namespace("hometicx"), zap.String("I'm", "Gopher"))

		c := context.WithValue(r.Context(), "logger", l)
		next.ServeHTTP(w, r.WithContext(c))
	})
}

func L(ctx context.Context) *zap.Logger {
	val := ctx.Value("logger")

	if val == nil {
		return zap.NewExample()
	}

	l, ok := val.(*zap.Logger)
	if ok {
		return l
	}
	return zap.NewExample()
}
