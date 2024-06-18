package infisical

import (
	"context"

	infisical "github.com/infisical/go-sdk"
	"github.com/infisical/go-sdk/packages/models"
)

type clientSecrets struct {
	client *Client
}

var _ infisical.SecretsInterface = clientSecrets{}

func (s clientSecrets) List(options infisical.ListSecretsOptions) ([]models.Secret, error) {
	c, err := s.client.client(context.Background())
	if err != nil {
		return nil, err
	}

	return c.Secrets().List(options)
}

func (s clientSecrets) Retrieve(options infisical.RetrieveSecretOptions) (models.Secret, error) {
	c, err := s.client.client(context.Background())
	if err != nil {
		return models.Secret{}, err
	}

	return c.Secrets().Retrieve(options)
}

func (s clientSecrets) Update(options infisical.UpdateSecretOptions) (models.Secret, error) {
	c, err := s.client.client(context.Background())
	if err != nil {
		return models.Secret{}, err
	}

	return c.Secrets().Update(options)
}

func (s clientSecrets) Create(options infisical.CreateSecretOptions) (models.Secret, error) {
	c, err := s.client.client(context.Background())
	if err != nil {
		return models.Secret{}, err
	}

	return c.Secrets().Create(options)
}

func (s clientSecrets) Delete(options infisical.DeleteSecretOptions) (models.Secret, error) {
	c, err := s.client.client(context.Background())
	if err != nil {
		return models.Secret{}, err
	}

	return c.Secrets().Delete(options)
}
