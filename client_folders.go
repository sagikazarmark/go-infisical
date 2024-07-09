package infisical

import (
	"context"

	infisical "github.com/infisical/go-sdk"
	"github.com/infisical/go-sdk/packages/models"
)

type clientFolders struct {
	client *Client
}

var _ FoldersInterface = clientFolders{}

type Folder = models.Folder

type ListFoldersOptions = infisical.ListFoldersOptions

func (f clientFolders) List(options ListFoldersOptions) ([]Folder, error) {
	c, err := f.client.client(context.Background())
	if err != nil {
		return nil, err
	}

	return c.Folders().List(options)
}

type UpdateFolderOptions = infisical.UpdateFolderOptions

func (f clientFolders) Update(options UpdateFolderOptions) (Folder, error) {
	c, err := f.client.client(context.Background())
	if err != nil {
		return models.Folder{}, err
	}

	return c.Folders().Update(options)
}

type CreateFolderOptions = infisical.CreateFolderOptions

func (f clientFolders) Create(options CreateFolderOptions) (Folder, error) {
	c, err := f.client.client(context.Background())
	if err != nil {
		return models.Folder{}, err
	}

	return c.Folders().Create(options)
}

type DeleteFolderOptions = infisical.DeleteFolderOptions

func (f clientFolders) Delete(options DeleteFolderOptions) (Folder, error) {
	c, err := f.client.client(context.Background())
	if err != nil {
		return models.Folder{}, err
	}

	return c.Folders().Delete(options)
}
