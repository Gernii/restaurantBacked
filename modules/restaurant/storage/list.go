package restaurantstorage

import (
	"context"
	"restaurantBacked/common"
	restaurantmodel "restaurantBacked/modules/restaurant/model"
)

func (s *sqlStore) ListDataWithCondition(
	ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
	moreKey ...string,
) ([]restaurantmodel.Restaurant, error) {
	db := s.db
	var result []restaurantmodel.Restaurant

	if filter.UserId > 0 {
		db = db.Where("owner_id = ?", filter.UserId)
	}
	db = db.Where("status not in (0)")
	if err := db.Table(restaurantmodel.Restaurant{}.TableName()).Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for i := range moreKey {
		db = db.Preload(moreKey[i])
	}

	offset := (paging.Page - 1) * paging.Limit

	if err := db.
		Limit(paging.Limit).
		Offset(offset).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return result, nil
}
