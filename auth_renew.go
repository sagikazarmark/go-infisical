package infisical

import (
	"context"
	"io"
	"log/slog"
	"sync"
	"time"

	"github.com/go-resty/resty/v2"
)

// AccessTokenRenewer wraps an [Authenticator] and renews the access token when it expires (either by [renewing an existing access token] (if possible) or issuing a new one).
//
// Note: this implementation uses a lock upon every token access. You may find it to be a bottleneck in high-throughput scenarios.
// An alternative implementation of this could refresh the token asynchronously and return the cached token in the meantime.
//
// [renewing an existing access token]: https://infisical.com/docs/api-reference/endpoints/universal-auth/renew-access-token
func AccessTokenRenewer(authenticator Authenticator, opts ...AccessTokenRenewerOption) Authenticator {
	a := &accessTokenRenewer{
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
type AccessTokenRenewerOption interface {
	apply(a *accessTokenRenewer)
}

type accessTokenRenewerOptionFunc func(a *accessTokenRenewer)

func (fn accessTokenRenewerOptionFunc) apply(a *accessTokenRenewer) {
	fn(a)
}

// WithExpirationWindow configures an expiration window for the access token.
// The access token will be refreshed when it expires or when the window is reached, whichever comes first.
func WithExpirationWindow(d time.Duration) AccessTokenRenewerOption {
	return accessTokenRenewerOptionFunc(func(a *accessTokenRenewer) {
		a.refreshWindow = d.Abs()
	})
}

// WithLogger configures a logger for the access token refresher.
func WithLogger(logger *slog.Logger) AccessTokenRenewerOption {
	return accessTokenRenewerOptionFunc(func(a *accessTokenRenewer) {
		a.logger = logger
	})
}

type accessTokenRenewer struct {
	authenticator Authenticator
	refreshWindow time.Duration

	credential *MachineIdentityCredential
	expiresAt  time.Time

	logger *slog.Logger

	mu sync.Mutex
}

var _ Authenticator = &accessTokenRenewer{}

func (a *accessTokenRenewer) Authenticate(ctx context.Context, httpClient *resty.Client) (MachineIdentityCredential, error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	// If we don't have a credential or it's about to expire, refresh it.
	if a.credential == nil || time.Now().After(a.expiresAt.Add(-a.refreshWindow)) {
		a.logger.Info("issuing new access token")

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
