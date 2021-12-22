package restaurantgin

import (
	"net/http"
	"restaurantBacked/common"
	"restaurantBacked/component/appctx"
	restaurantbiz "restaurantBacked/modules/restaurant/biz"
	restaurantstorage "restaurantBacked/modules/restaurant/storage"

	"github.com/gin-gonic/gin"
)

func DeleteRestaurant(appCtx appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		requester := c.MustGet(common.CurrentUser).(common.Requester)

		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewDeleteRestaurantBiz(store, requester)

		if err := biz.DeleteRestaurant(c.Request.Context(), int(uid.GetLocalID())); err != nil {
			panic(err)

		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
