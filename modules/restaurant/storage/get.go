package restaurantstorage

import (
	"context"
	"restaurantBacked/common"
	restaurantmodel "restaurantBacked/modules/restaurant/model"

	"gorm.io/gorm"
)

func (s *sqlStore) FindDataWithCondition(
	ctx context.Context,
	cond map[string]interface{},
	moreKey ...string,
) (*restaurantmodel.Restaurant, error) {
	db := s.db
	var data restaurantmodel.Restaurant

	for i := range moreKey {
		db = db.Preload(moreKey[i])
	}

	if err := db.Where(cond).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrDataNotFound
		}
		return nil, common.ErrDB(err)
	}
	return &data, nil
}
