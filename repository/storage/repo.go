package repository

import "github.com/xxxibgdrgnmm/reverse-registry/model"

type Interface interface {
	FindByNameTag(nameWithTag string) (*model.ImageModel, error)
	FindByDigest(digest string) (*model.ImageModel, error)
	SaveDigest(nameWithTag, digest string) error
}
