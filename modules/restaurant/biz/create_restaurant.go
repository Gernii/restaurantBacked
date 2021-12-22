package restaurantbiz

import (
	"context"
	"restaurantBacked/common"
	restaurantmodel "restaurantBacked/modules/restaurant/model"
)

type CreateRestaurantStore interface {
	Create(ctx context.Context, data *restaurantmodel.RestaurantCreate) error
}

type createNewRestaurantBiz struct {
	store     CreateRestaurantStore
	requester common.Requester
}

func NewCreateRestaurantBiz(store CreateRestaurantStore, requester common.Requester) *createNewRestaurantBiz {
	return &createNewRestaurantBiz{store: store, requester: requester}
}

func (biz *createNewRestaurantBiz) CreateNewRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
	if err := data.Validate(); err != nil {
		return common.ErrInvalidRequest(err)
	}

	if err := biz.store.Create(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(restaurantmodel.EntityName, err)
	}
	return nil
}
