package usergin

import (
	"net/http"
	"restaurantBacked/common"
	"restaurantBacked/component/appctx"
	"restaurantBacked/component/hasher"
	"restaurantBacked/component/tokenprovider/jwt"
	userbiz "restaurantBacked/modules/user/biz"
	usermodel "restaurantBacked/modules/user/model"
	userstorage "restaurantBacked/modules/user/storage"

	"github.com/gin-gonic/gin"
)

func Login(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginUserData usermodel.UserLogin
		if err := c.ShouldBind(&loginUserData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		db := appCtx.GetMainDBConnection()
		tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())

		store := userstorage.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()
		business := userbiz.NewLoginBusiness(store, tokenProvider, md5, 60*60*24*30)
		account, err := business.Login(c.Request.Context(), &loginUserData)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(account))
	}
}
