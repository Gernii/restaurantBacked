package restaurantbiz

import (
	"context"
	"restaurantBacked/common"
	restaurantmodel "restaurantBacked/modules/restaurant/model"
)

type ListRestaurantStore interface {
	ListDataWithCondition(
		ctx context.Context,
		filter *restaurantmodel.Filter,
		paging *common.Paging,
		moreKey ...string,
	) ([]restaurantmodel.Restaurant, error)
}

type listNewRestaurantBiz struct {
	store     ListRestaurantStore
	requester common.Requester
}

func NewListRestaurantBiz(store ListRestaurantStore, requester common.Requester) *listNewRestaurantBiz {
	return &listNewRestaurantBiz{store: store, requester: requester}
}

func (biz *listNewRestaurantBiz) ListRestaurant(
	ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
) ([]restaurantmodel.Restaurant, error) {

	result, err := biz.store.ListDataWithCondition(ctx, filter, paging, "User")

	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantmodel.EntityName, err)
	}
	return result, nil
}
