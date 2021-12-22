package restaurantstorage

import (
	"context"
	"restaurantBacked/common"
	restaurantmodel "restaurantBacked/modules/restaurant/model"
)

func (s *sqlStore) Update(
	ctx context.Context,
	cond map[string]interface{},
	updateData *restaurantmodel.RestaurantUpdate,
) error {
	db := s.db
	if err := db.Where(cond).Updates(updateData).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
