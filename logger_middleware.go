package web

import (
	"fmt"
	"time"
)

func LoggerMiddleware(rw ResponseWriter, req *Request, next NextMiddlewareFunc) error {
	startTime := time.Now()

	if err := next(rw, req); err != nil {
		return err
	}

	duration := time.Since(startTime).Nanoseconds()
	var durationUnits string
	switch {
	case duration > 2000000:
		durationUnits = "ms"
		duration /= 1000000
	case duration > 1000:
		durationUnits = "μs"
		duration /= 1000
	default:
		durationUnits = "ns"
	}

	fmt.Printf("[%d %s] %d '%s'\n", duration, durationUnits, rw.StatusCode(), req.URL.Path)
	return nil
}
