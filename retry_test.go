package retry

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMultipleTimes(t *testing.T) {
	i := 0
	err := MultipleTimes(3, func() error {
		i++
		if i == 4 {
			return nil
		}

		return fmt.Errorf("error")
	})

	require.NoError(t, err)

	// 1 failed run + 3 retries
	require.Equal(t, 4, i)
}

func TestMultipleTimesWithDelay(t *testing.T) {
	i := 0
	err := MultipleTimes(2, func() error {
		i++
		if i == 3 {
			return nil
		}

		return fmt.Errorf("error")
	})

	require.NoError(t, err)

	// 1 failed run + 2 retries
	require.Equal(t, 3, i)
}

func TestUntilError(t *testing.T) {
	i := 0
	err := UntilError(func() error {
		i++
		if i == 3 {
			return fmt.Errorf("error")
		}

		return nil
	})

	require.Error(t, err)

	// 1 failed run + 2 retries
	require.Equal(t, 3, i)
}

func TestUntilErrorOrCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	i := 0
	err := UntilErrorOrCancel(ctx, func() error {
		i++
		if i == 3 {
			cancel()
			return nil
		}

		return nil
	})

	require.ErrorIs(t, err, context.Canceled)

	// 1 failed run + 2 retries
	require.Equal(t, 3, i)
}

func TestUntilSuccess(t *testing.T) {
	i := 0
	err := UntilSuccess(func() error {
		i++
		if i == 3 {
			return nil
		}

		return fmt.Errorf("error")
	})

	require.NoError(t, err)

	// 1 failed run + 2 retries
	require.Equal(t, 3, i)
}

func TestUntilSuccessOrCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	i := 0
	err := UntilSuccessOrCancel(ctx, func() error {
		i++
		if i == 3 {
			cancel()
			return fmt.Errorf("error")
		}

		return fmt.Errorf("error")
	})

	require.ErrorIs(t, err, context.Canceled)

	// 1 failed run + 2 retries
	require.Equal(t, 3, i)
}
