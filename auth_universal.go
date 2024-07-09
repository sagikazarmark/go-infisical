package infisical

import (
	"context"

	"github.com/go-resty/resty/v2"
	api "github.com/infisical/go-sdk/packages/api/auth"
)

// UniversalAuthLogin authenticates using a client ID and client secret.
func UniversalAuthLogin(clientID string, clientSecret string) Authenticator {
	return universalAuthenticator{
		clientID:     clientID,
		clientSecret: clientSecret,
	}
}

type universalAuthenticator struct {
	clientID     string
	clientSecret string
}

var _ Authenticator = universalAuthenticator{}

func (a universalAuthenticator) Authenticate(_ context.Context, httpClient *resty.Client) (MachineIdentityCredential, error) {
	return api.CallUniversalAuthLogin(httpClient, api.UniversalAuthLoginRequest{
		ClientID:     a.clientID,
		ClientSecret: a.clientSecret,
	})
}
