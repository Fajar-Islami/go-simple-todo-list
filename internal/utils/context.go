package utils

import (
	"context"
	"os"
	"strconv"
	"time"
)

func ContextWithTimeout(ctx context.Context) (context.Context, context.CancelFunc) {
	timeout := 60

	if os.Getenv("CONTEXT_TIMEOUT") != "" {
		timeout, _ = strconv.Atoi(os.Getenv("CONTEXT_TIMEOUT"))
	}

	return context.WithTimeout(ctx, time.Duration(timeout)*time.Second)
}

func ContextWithCustomTimeout(ctx context.Context, timeout int) (context.Context, context.CancelFunc) {
	return context.WithTimeout(ctx, time.Duration(timeout)*time.Second)
}
