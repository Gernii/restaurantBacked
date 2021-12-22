package restaurantgin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"restaurantBacked/common"
	"restaurantBacked/component/appctx"
	restaurantbiz "restaurantBacked/modules/restaurant/biz"
	restaurantstorage "restaurantBacked/modules/restaurant/storage"
)

func GetRestaurant(appCtx appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))

		}
		requester := c.MustGet(common.CurrentUser).(common.Requester)

		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewGetRestaurantBiz(store, requester)

		data, err := biz.GetRestaurant(c.Request.Context(), int(uid.GetLocalID()))

		if err != nil {
			panic(common.ErrInvalidRequest(err))

		}

		data.Mask(common.DbTypeRestaurant)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
