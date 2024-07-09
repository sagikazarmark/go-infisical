package infisical

import (
	"context"

	"github.com/go-resty/resty/v2"
	infisical "github.com/infisical/go-sdk"
)

type MachineIdentityCredential = infisical.MachineIdentityCredential

// Authenticator exchanges a set of credentials (either explicit or identity-based) for an access token.
type Authenticator interface {
	Authenticate(ctx context.Context, httpClient *resty.Client) (MachineIdentityCredential, error)
}
