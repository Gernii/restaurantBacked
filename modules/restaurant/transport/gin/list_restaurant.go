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

func ListRestaurant(appCtx appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var filter restaurantmodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))

		}

		if err := paging.Process(); err != nil {
			panic(common.ErrInvalidRequest(err))

		}
		requester := c.MustGet(common.CurrentUser).(common.Requester)

		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewListRestaurantBiz(store, requester)

		result, err := biz.ListRestaurant(c.Request.Context(), &filter, &paging)
		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(common.DbTypeRestaurant)

		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
