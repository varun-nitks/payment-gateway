package internal

import (
	"time"
)

func Retry(attempts int, sleep time.Duration, fn func() error) error {
	for i := 0; i < attempts; i++ {
		if err := fn(); err == nil {
			return nil
		}
		time.Sleep(sleep)
	}
	return nil
}
