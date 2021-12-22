package uploadprovider

import (
	"context"
	"restaurantBacked/common"
)

type UploadProvider interface {
	SaveFileUploaded(ctx context.Context, data []byte, dst string) (*common.Image, error)
	GetDomain() string
}
