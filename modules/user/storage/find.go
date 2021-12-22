package userstorage

import (
	"context"
	"restaurantBacked/common"
	usermodel "restaurantBacked/modules/user/model"

	"gorm.io/gorm"
)

func (s *sqlStore) FindUser(ctx context.Context, conditions map[string]interface{}, moreinfo ...string) (*usermodel.User, error) {
	db := s.db.Table(usermodel.User{}.TableName())
	for i := range moreinfo {
		db = db.Preload(moreinfo[i])
	}
	var user usermodel.User
	// _, span := trace.StartSpan(ctx, "store.user.find_user")

	if err := db.Where(conditions).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}
	return &user, nil
}
