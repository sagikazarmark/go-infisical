package infisical

import (
	"context"
	"io"
	"log/slog"
	"sync"
	"time"

	"github.com/go-resty/resty/v2"
)

// AccessTokenRefresher wraps an [Authenticator] and refreshes the access token when it expires.
//
// Note: this implementation uses a lock upon every token access. You may find it to be a bottleneck in high-throughput scenarios.
// An alternative implementation of this could refresh the token asynchronously and return the cached token in the meantime.
func AccessTokenRefresher(authenticator Authenticator, opts ...AccessTokenRefresherOption) Authenticator {
	a := &accessTokenRefresher{
		authenticator: authenticator,
	}

	for _, opt := range opts {
		opt.apply(a)
	}

	if a.logger == nil {
		a.logger = slog.New(slog.NewTextHandler(io.Discard, nil))
	}

	return a
}

// Option configures a [Client] using the functional options paradigm popularized by Rob Pike and Dave Cheney.
// If you're unfamiliar with this style,
// see https://commandcenter.blogspot.com/2014/01/self-referential-functions-and-design.html and
// https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis.
type AccessTokenRefresherOption interface {
	apply(a *accessTokenRefresher)
}

type accessTokenRefresherOptionFunc func(a *accessTokenRefresher)

func (fn accessTokenRefresherOptionFunc) apply(a *accessTokenRefresher) {
	fn(a)
}

// WithRefreshWindow configures a refresh window for the access token.
// The access token will be refreshed when it expires or when the refresh window is reached, whichever comes first.
func WithRefreshWindow(d time.Duration) AccessTokenRefresherOption {
	return accessTokenRefresherOptionFunc(func(a *accessTokenRefresher) {
		a.refreshWindow = d.Abs()
	})
}

// WithLogger configures a logger for the access token refresher.
func WithLogger(logger *slog.Logger) AccessTokenRefresherOption {
	return accessTokenRefresherOptionFunc(func(a *accessTokenRefresher) {
		a.logger = logger
	})
}

type accessTokenRefresher struct {
	authenticator Authenticator
	refreshWindow time.Duration

	credential *MachineIdentityCredential
	expiresAt  time.Time

	logger *slog.Logger

	mu sync.Mutex
}

var _ Authenticator = &accessTokenRefresher{}

func (a *accessTokenRefresher) Authenticate(ctx context.Context, httpClient *resty.Client) (MachineIdentityCredential, error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	// If we don't have a credential or it's about to expire, refresh it.
	if a.credential == nil || time.Now().After(a.expiresAt.Add(-a.refreshWindow)) {
		a.logger.Info("refreshing access token")

		credential, err := a.authenticator.Authenticate(ctx, httpClient)
		if err != nil {
			return MachineIdentityCredential{}, err
		}

		a.credential = &credential
		a.expiresAt = time.Now().Add(time.Second * time.Duration(credential.ExpiresIn))

		a.logger.Debug("new access token saved", slog.Time("expiresAt", a.expiresAt))

		// Just fetched a new credential, we can safely assume it's valid.
		return credential, nil
	}

	a.logger.Debug("returning cached access token", slog.Time("expiresAt", a.expiresAt))

	return *a.credential, nil
}
