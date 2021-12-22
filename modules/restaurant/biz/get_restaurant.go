package restaurantbiz

import (
	"context"
	"restaurantBacked/common"
	restaurantmodel "restaurantBacked/modules/restaurant/model"
)

type GetRestaurantStore interface {
	FindDataWithCondition(
		ctx context.Context,
		cond map[string]interface{},
		moreKey ...string,
	) (*restaurantmodel.Restaurant, error)
}

type getRestaurantBiz struct {
	store     GetRestaurantStore
	requester common.Requester
}

func NewGetRestaurantBiz(store GetRestaurantStore, requester common.Requester) *getRestaurantBiz {
	return &getRestaurantBiz{store: store, requester: requester}
}

func (biz *getRestaurantBiz) GetRestaurant(ctx context.Context, id int) (*restaurantmodel.Restaurant, error) {
	result, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{"id": id}, "User")
	if err != nil {
		return nil, err
	}
	return result, nil
}
