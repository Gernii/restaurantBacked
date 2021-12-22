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

func CreateRestaurant(appCtx appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		var newData restaurantmodel.RestaurantCreate
		if err := c.ShouldBind(&newData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		// Dependencies install
		requester := c.MustGet(common.CurrentUser).(common.Requester)
		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewCreateRestaurantBiz(store, requester)

		if err := biz.CreateNewRestaurant(c.Request.Context(), &newData); err != nil {
			panic(err)
		}

		newData.Mask(common.DbTypeRestaurant)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(newData.FakeId))
	}
}
