package main

import (
	"github.com/gin-gonic/gin"
	"restaurantBacked/component/appctx"
	"restaurantBacked/middleware"
	restaurantgin "restaurantBacked/modules/restaurant/transport/gin"
	userstorage "restaurantBacked/modules/user/storage"
)

func mainRoute(g *gin.RouterGroup, appCtx appctx.AppContext) {
	authStore := userstorage.NewSQLStore(appCtx.GetMainDBConnection())

	restaurants := g.Group("/restaurants", middleware.RequiredAuth(appCtx, authStore))
	{
		restaurants.POST("", restaurantgin.CreateRestaurant(appCtx))
		restaurants.GET("/:id", restaurantgin.GetRestaurant(appCtx))
		restaurants.GET("", restaurantgin.ListRestaurant(appCtx))
		restaurants.PUT("/:id", restaurantgin.UpdateRestaurant(appCtx))
		restaurants.DELETE("/:id", restaurantgin.DeleteRestaurant(appCtx))

	}
}
