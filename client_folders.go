package infisical

import (
	"context"

	infisical "github.com/infisical/go-sdk"
	"github.com/infisical/go-sdk/packages/models"
)

type clientFolders struct {
	client *Client
}

var _ infisical.FoldersInterface = clientFolders{}

func (f clientFolders) List(options infisical.ListFoldersOptions) ([]models.Folder, error) {
	c, err := f.client.client(context.Background())
	if err != nil {
		return nil, err
	}

	return c.Folders().List(options)
}

func (f clientFolders) Update(options infisical.UpdateFolderOptions) (models.Folder, error) {
	c, err := f.client.client(context.Background())
	if err != nil {
		return models.Folder{}, err
	}

	return c.Folders().Update(options)
}

func (f clientFolders) Create(options infisical.CreateFolderOptions) (models.Folder, error) {
	c, err := f.client.client(context.Background())
	if err != nil {
		return models.Folder{}, err
	}

	return c.Folders().Create(options)
}

func (f clientFolders) Delete(options infisical.DeleteFolderOptions) (models.Folder, error) {
	c, err := f.client.client(context.Background())
	if err != nil {
		return models.Folder{}, err
	}

	return c.Folders().Delete(options)
}
