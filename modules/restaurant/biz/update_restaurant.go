package restaurantbiz

import (
	"context"
	"errors"
	"fmt"
	"restaurantBacked/common"
	restaurantmodel "restaurantBacked/modules/restaurant/model"
)

type UpdateRestaurantStore interface {
	FindDataWithCondition(
		ctx context.Context,
		cond map[string]interface{},
		moreKey ...string,
	) (*restaurantmodel.Restaurant, error)
	Update(
		ctx context.Context,
		cond map[string]interface{},
		updateData *restaurantmodel.RestaurantUpdate,
	) error
}

type updateRestaurantBiz struct {
	store     UpdateRestaurantStore
	requester common.Requester
}

func NewUpdateRestaurantBiz(store UpdateRestaurantStore, requester common.Requester) *updateRestaurantBiz {
	return &updateRestaurantBiz{store: store, requester: requester}
}

func (biz *updateRestaurantBiz) UpdateRestaurant(
	ctx context.Context,
	id int,
	data *restaurantmodel.RestaurantUpdate,
) error {
	if err := data.Validate(); err != nil {
		return err
	}
	oldData, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		if err == common.ErrDataNotFound {
			return errors.New("data not found")
		}
		return err
	}
	fmt.Println(oldData)
	fmt.Println(oldData.Status)
	if oldData.Status == 0 {
		return errors.New("data has been deleted")
	}

	if biz.requester.GetRole() != "admin" && oldData.OwnerId != biz.requester.GetUserId() {
		return common.ErrNoPermission(nil)
	}

	if err := biz.store.Update(ctx, map[string]interface{}{"id": id}, data); err != nil {
		return err
	}
	return nil
}
