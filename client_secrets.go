package infisical

import (
	"context"

	infisical "github.com/infisical/go-sdk"
	"github.com/infisical/go-sdk/packages/models"
)

type clientSecrets struct {
	client *Client
}

var _ SecretsInterface = clientSecrets{}

type Secret = models.Secret

type ListSecretsOptions = infisical.ListSecretsOptions

func (s clientSecrets) List(options ListSecretsOptions) ([]Secret, error) {
	c, err := s.client.client(context.Background())
	if err != nil {
		return nil, err
	}

	return c.Secrets().List(options)
}

type RetrieveSecretOptions = infisical.RetrieveSecretOptions

func (s clientSecrets) Retrieve(options RetrieveSecretOptions) (Secret, error) {
	c, err := s.client.client(context.Background())
	if err != nil {
		return models.Secret{}, err
	}

	return c.Secrets().Retrieve(options)
}

type UpdateSecretOptions = infisical.UpdateSecretOptions

func (s clientSecrets) Update(options UpdateSecretOptions) (Secret, error) {
	c, err := s.client.client(context.Background())
	if err != nil {
		return models.Secret{}, err
	}

	return c.Secrets().Update(options)
}

type CreateSecretOptions = infisical.CreateSecretOptions

func (s clientSecrets) Create(options CreateSecretOptions) (Secret, error) {
	c, err := s.client.client(context.Background())
	if err != nil {
		return models.Secret{}, err
	}

	return c.Secrets().Create(options)
}

type DeleteSecretOptions = infisical.DeleteSecretOptions

func (s clientSecrets) Delete(options DeleteSecretOptions) (Secret, error) {
	c, err := s.client.client(context.Background())
	if err != nil {
		return models.Secret{}, err
	}

	return c.Secrets().Delete(options)
}
