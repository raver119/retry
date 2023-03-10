package retry

import (
	"context"
	"time"
)

// UntilError retries the given function until it returns an error
func UntilError(f func() error) error {
	for {
		err := f()
		if err != nil {
			return err
		}
	}
}

// UntilErrorWithDelay retries the given function until it returns an error, sleeping between each attempt
func UntilErrorWithDelay(delay time.Duration, f func() error) error {
	for {
		if err := f(); err != nil {
			return err
		}

		time.Sleep(delay)
	}
}

// UntilErrorOrCancel retries the given function until it returns an error or the context is canceled
func UntilErrorOrCancel(ctx context.Context, f func() error) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			if err := f(); err != nil {
				return err
			}
		}
	}
}

// UntilErrorOrTimeout retries the given function until it returns an error or the timeout is reached
func UntilErrorOrTimeout(timeout time.Duration, f func() error) error {
	start := time.Now()
	for {
		if err := f(); err != nil {
			return err
		}

		if time.Since(start) > timeout {
			return ErrTimeout
		}
	}
}

// UntilSuccess retries the given function until it returns true
func UntilSuccess(f func() error) error {
	for {
		if err := f(); err == nil {
			return nil
		}
	}
}

// UntilSuccessWithDelay retries the given function until it returns true, sleep, or the timeout is reached
func UntilSuccessWithDelay(delay time.Duration, f func() error) error {
	for {
		if err := f(); err == nil {
			return nil
		}

		time.Sleep(delay)
	}
}

// UntilSuccessOrTimeout retries the given function until it returns true or the timeout is reached
func UntilSuccessOrTimeout(ctx context.Context, timeout time.Duration, f func() error) error {
	start := time.Now()
	for {
		if err := f(); err == nil {
			return nil
		}

		if time.Since(start) > timeout {
			return ErrTimeout
		}
	}
}

// UntilSuccessOrCancel retries the given function until it returns true or the context is canceled
func UntilSuccessOrCancel(ctx context.Context, f func() error) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			if err := f(); err == nil {
				return nil
			}
		}
	}
}

// MultipleTimes retries the given function multiple times, or until it returns no error
func MultipleTimes(times int, f func() error) (err error) {
	for i := 0; i <= times; i++ {
		if err = f(); err == nil {
			return nil
		}
	}

	// return the last error
	return err
}

// MultipleTimesWithDelay retries the given function multiple times, sleep, or until it returns no error
func MultipleTimesWithDelay(times int, delay time.Duration, f func() error) (err error) {
	for i := 0; i <= times; i++ {
		if err = f(); err == nil {
			return nil
		}

		// wait for the delay, if we're not on the last iteration
		if i < times-1 && delay > 0 {
			time.Sleep(delay)
		}
	}

	// return the last error
	return err
}

// Once runs a given function and retries it once if it returns an error
func Once(f func() error) error {
	return MultipleTimes(1, f)
}

// OnceWithSmallDelay runs a given function and retries it once if it returns an error, sleeping before the retry
func OnceWithSmallDelay(f func() error) error {
	return MultipleTimesWithDelay(1, 100*time.Millisecond, f)
}

// Twice runs a given function and retries it twice if it returns an error
func Twice(f func() error) error {
	return MultipleTimes(2, f)
}

// TwiceWithSmallDelay runs a given function and retries it twice if it returns an error, sleeping before the retries
func TwiceWithSmallDelay(f func() error) error {
	return MultipleTimesWithDelay(2, 100*time.Millisecond, f)
}

// Thrice runs a given function and retries it three times if it returns an error
func Thrice(f func() error) error {
	return MultipleTimes(3, f)
}

// ThriceWithSmallDelay runs a given function and retries it three times if it returns an error, sleeping before the retries
func ThriceWithSmallDelay(f func() error) error {
	return MultipleTimesWithDelay(3, 100*time.Millisecond, f)
}

// UntilSuccededOrCancelled function will retry the given function until it returns no error or is cancelled via the context.
func UntilSucceededOrCancelledWithDelay(ctx context.Context, delay time.Duration, f func() error) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			if err := f(); err == nil {
				return nil
			}

			// sleep for the given delay, if set
			if delay > 0 {
				time.Sleep(delay)
			}
		}
	}
}
