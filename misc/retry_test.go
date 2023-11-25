package misc

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestRetry(t *testing.T) {
	t.Run("function succeeds on first try", func(t *testing.T) {
		f := func() error {
			return nil
		}
		err := Retry(context.Background(), f, 5, time.Millisecond)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
	})

	t.Run("function fails on all tries", func(t *testing.T) {
		f := func() error {
			return errors.New("error")
		}
		err := Retry(context.Background(), f, 5, time.Millisecond)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})

	t.Run("function succeeds after a few tries", func(t *testing.T) {
		var callCount int
		f := func() error {
			callCount++
			if callCount < 3 {
				return errors.New("error")
			}
			return nil
		}
		err := Retry(context.Background(), f, 5, time.Millisecond)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
	})

	t.Run("context is cancelled before function can succeed", func(t *testing.T) {
		var callCount int
		f := func() error {
			callCount++
			if callCount < 3 {
				return errors.New("error")
			}
			return nil
		}
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		err := Retry(ctx, f, 5, time.Millisecond)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})
}
