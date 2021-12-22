package restaurantgin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"restaurantBacked/common"
	"restaurantBacked/component/appctx"
	restaurantbiz "restaurantBacked/modules/restaurant/biz"
	restaurantmodel "restaurantBacked/modules/restaurant/model"
	restaurantstorage "restaurantBacked/modules/restaurant/storage"
)

func UpdateRestaurant(appCtx appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		id, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		var data restaurantmodel.RestaurantUpdate
		requester := c.MustGet(common.CurrentUser).(common.Requester)

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewUpdateRestaurantBiz(store, requester)

		if err := biz.UpdateRestaurant(c.Request.Context(), int(id.GetLocalID()), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
