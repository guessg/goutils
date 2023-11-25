package misc

import (
	"context"
	"time"
)

func Retry(ctx context.Context, f func() error, times int, sleepInterval time.Duration) error {
	var err error
	for i := 0; i < times; i++ {
		err = f()
		if err == nil {
			return nil
		}
		select {
		case <-ctx.Done():
			return err
		case <-time.After(sleepInterval):
		}
	}
	return err
}
