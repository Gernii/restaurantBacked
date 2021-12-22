package appctx

import (
	"restaurantBacked/component/uploadprovider"

	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	UploadProvider() uploadprovider.UploadProvider
	SecretKey() string
}

type appContext struct {
	db             *gorm.DB
	uploadProvider uploadprovider.UploadProvider
	secret         string
}

func NewAppContext(db *gorm.DB, uploadProvider uploadprovider.UploadProvider, secret string) *appContext {
	return &appContext{db: db, uploadProvider: uploadProvider, secret: secret}
}

func (appCtx *appContext) GetMainDBConnection() *gorm.DB {
	return appCtx.db
}

func (appCtx *appContext) UploadProvider() uploadprovider.UploadProvider {
	return appCtx.uploadProvider
}

func (appCtx *appContext) SecretKey() string {
	return appCtx.secret
}
