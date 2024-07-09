package infisical

import (
	"context"

	"github.com/go-resty/resty/v2"
	infisical "github.com/infisical/go-sdk"
	"github.com/infisical/go-sdk/packages/util"
)

//nolint:revive
type InfisicalClientInterface = infisical.InfisicalClientInterface

type Client struct {
	authenticator Authenticator

	siteURL   string
	userAgent string

	// necessary for authenticators to work
	httpClient *resty.Client
}

var _ infisical.InfisicalClientInterface = &Client{}

// NewClient returns a new [Client].
func NewClient(authenticator Authenticator, opts ...ClientOption) *Client {
	client := &Client{
		authenticator: authenticator,
	}

	for _, opt := range opts {
		opt.apply(client)
	}

	// TODO: introduce an options struct so we don't store these on the client
	siteURL := client.siteURL
	if siteURL == "" {
		siteURL = util.DEFAULT_INFISICAL_API_URL
	}
	siteURL = util.AppendAPIEndpoint(siteURL)

	userAgent := client.userAgent
	if userAgent == "" {
		userAgent = "infisical-go-sdk"
	}

	client.httpClient = resty.New().SetHeader("User-Agent", userAgent).SetBaseURL(siteURL)

	return client
}

// ClientOption configures a [Client] using the functional options paradigm popularized by Rob Pike and Dave Cheney.
// If you're unfamiliar with this style,
// see https://commandcenter.blogspot.com/2014/01/self-referential-functions-and-design.html and
// https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis.
type ClientOption interface {
	apply(c *Client)
}

type clientOptionFunc func(c *Client)

func (fn clientOptionFunc) apply(c *Client) {
	fn(c)
}

// WithSiteURL configures a custom site URL.
func WithSiteURL(s string) ClientOption {
	return clientOptionFunc(func(c *Client) {
		c.siteURL = s
	})
}

// WithUserAgent configures a custom user agent.
func WithUserAgent(s string) ClientOption {
	return clientOptionFunc(func(c *Client) {
		c.userAgent = s
	})
}

func (c *Client) client(ctx context.Context) (InfisicalClientInterface, error) {
	config := infisical.Config{
		SiteUrl:   c.siteURL,
		UserAgent: c.userAgent,
	}

	client := infisical.NewInfisicalClient(config)

	credential, err := c.authenticator.Authenticate(ctx, c.httpClient)
	if err != nil {
		return nil, err
	}

	client.Auth().SetAccessToken(credential.AccessToken)

	return client, nil
}

type Config = infisical.Config

// UpdateConfiguration implements [InfisicalClientInterface].
// DO NOT USE THIS FUNCTION: IT PANICS.
//
// Updating configuration after initialization is bad practice.
// Initialize a new client instead.
func (c *Client) UpdateConfiguration(_ infisical.Config) {
	panic("do not update configuration after initialization: initialize a new client instead")
}

type SecretsInterface = infisical.SecretsInterface

// Secrets implements [InfisicalClientInterface].
func (c *Client) Secrets() SecretsInterface {
	return clientSecrets{c}
}

type FoldersInterface = infisical.FoldersInterface

// Folders implements [InfisicalClientInterface].
func (c *Client) Folders() FoldersInterface {
	return clientFolders{c}
}

type AuthInterface = infisical.AuthInterface

// Auth implements [InfisicalClientInterface].
//
// Note: Calling any of the auth functions will not change the access token or authentication method of the client.
func (c *Client) Auth() AuthInterface {
	return clientAuth{c}
}
