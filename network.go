package retry

import (
	"context"
	"fmt"
	"net"
	"time"
)

// ConnectionUntilConnected retries the given address:port until a TCP connection established
func ConnectionUntilConnected(address string, port int) error {
	return UntilSuccess(func() error {
		_, err := net.Dial("tcp", fmt.Sprintf("%s:%d", address, port))
		return err
	})
}

// ConnectionUntilConnectedOrTimeout retries the given address:port until a TCP connection established or the timeout is reached
func ConnectionUntilConnectedOrTimeout(timeout time.Duration, address string, port int) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	return UntilSuccessOrCancel(ctx, func() error {
		_, err := net.Dial("tcp", fmt.Sprintf("%s:%d", address, port))
		return err
	})
}
