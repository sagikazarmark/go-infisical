package infisical

import (
	"context"

	infisical "github.com/infisical/go-sdk"
)

type clientAuth struct {
	client *Client
}

var _ infisical.AuthInterface = clientAuth{}

func (a clientAuth) SetAccessToken(_ string) {
	// This function call doesn't make sense in the context of this implementation
}

func (a clientAuth) UniversalAuthLogin(clientID string, clientSecret string) (credential infisical.MachineIdentityCredential, err error) {
	c, err := a.client.client(context.Background())
	if err != nil {
		return infisical.MachineIdentityCredential{}, err
	}

	return c.Auth().UniversalAuthLogin(clientID, clientSecret)
}

func (a clientAuth) KubernetesAuthLogin(identityID string, serviceAccountTokenPath string) (credential infisical.MachineIdentityCredential, err error) {
	c, err := a.client.client(context.Background())
	if err != nil {
		return infisical.MachineIdentityCredential{}, err
	}

	return c.Auth().KubernetesAuthLogin(identityID, serviceAccountTokenPath)
}

func (a clientAuth) KubernetesRawServiceAccountTokenLogin(identityID string, serviceAccountToken string) (credential infisical.MachineIdentityCredential, err error) {
	c, err := a.client.client(context.Background())
	if err != nil {
		return infisical.MachineIdentityCredential{}, err
	}

	return c.Auth().KubernetesRawServiceAccountTokenLogin(identityID, serviceAccountToken)
}

func (a clientAuth) AzureAuthLogin(identityID string) (credential infisical.MachineIdentityCredential, err error) {
	c, err := a.client.client(context.Background())
	if err != nil {
		return infisical.MachineIdentityCredential{}, err
	}

	return c.Auth().AzureAuthLogin(identityID)
}

func (a clientAuth) GcpIdTokenAuthLogin(identityID string) (credential infisical.MachineIdentityCredential, err error) {
	c, err := a.client.client(context.Background())
	if err != nil {
		return infisical.MachineIdentityCredential{}, err
	}

	return c.Auth().GcpIdTokenAuthLogin(identityID)
}

func (a clientAuth) GcpIamAuthLogin(identityID string, serviceAccountKeyFilePath string) (credential infisical.MachineIdentityCredential, err error) {
	c, err := a.client.client(context.Background())
	if err != nil {
		return infisical.MachineIdentityCredential{}, err
	}

	return c.Auth().GcpIamAuthLogin(identityID, serviceAccountKeyFilePath)
}

func (a clientAuth) AwsIamAuthLogin(identityId string) (credential infisical.MachineIdentityCredential, err error) {
	c, err := a.client.client(context.Background())
	if err != nil {
		return infisical.MachineIdentityCredential{}, err
	}

	return c.Auth().AwsIamAuthLogin(identityId)
}
