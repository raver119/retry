package retry

import (
	"net/http"
	"net/http/httptest"
	"net/netip"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestConnectionUntilConnected(t *testing.T) {
	require.Error(t, ConnectionUntilConnectedOrTimeout(time.Second, "localhost", 17821))

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
		return
	}))
	defer srv.Close()

	ap, err := netip.ParseAddrPort(srv.Listener.Addr().String())
	require.NoError(t, err)
	require.NotZero(t, ap.Port())

	require.NoError(t, ConnectionUntilConnectedOrTimeout(time.Second, "localhost", int(ap.Port())))
}
