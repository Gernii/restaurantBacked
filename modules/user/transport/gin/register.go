package usergin

import (
	"net/http"
	"restaurantBacked/common"
	"restaurantBacked/component/appctx"
	"restaurantBacked/component/hasher"
	userbiz "restaurantBacked/modules/user/biz"
	usermodel "restaurantBacked/modules/user/model"
	userstorage "restaurantBacked/modules/user/storage"

	"github.com/gin-gonic/gin"
)

func Register(appCtx appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		var data usermodel.UserCreate
		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := userstorage.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()
		biz := userbiz.NewRegisterBusiness(store, md5)

		if err := biz.Register(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		data.Mask(common.DbTypeUser)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}

}
