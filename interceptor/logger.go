package interceptor

import (
	"context"
	"log/slog"

	"connectrpc.com/connect"
)

// Logger is a unary interceptor that logs requests and responses.
func Logger() connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			slog.Info("request", slog.Any("request", req))
			return next(ctx, req)
		}
	}
}
