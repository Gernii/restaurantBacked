package middleware

import (
	"github.com/gin-gonic/gin"
	"restaurantBacked/common"
	"restaurantBacked/component/appctx"
)

func RequiredRoles(appCtx appctx.AppContext, roles ...string) func(*gin.Context) {
	return func(c *gin.Context) {
		u := c.MustGet(common.CurrentUser).(common.Requester)

		for i := range roles {
			if u.GetRole() == roles[i] {
				c.Next()
				return
			}
		}

		panic(common.ErrNoPermission(nil))
	}

}
