package restaurantmodel

import (
	"restaurantBacked/common"
	"strings"
)

type RestaurantUpdate struct {
	Name    *string        `json:"name" gorm:"column:name;"`
	Address *string        `json:"address" gorm:"column:addr;"`
	OwnerId int            `json:"-" gorm:"column:owner_id;"`
	Status  *int           `json:"_" gorm:"status;"`
	Cover   *common.Images `json:"cover" gorm:"column:cover;"`
}

func (RestaurantUpdate) TableName() string { return Restaurant{}.TableName() }

func (u *RestaurantUpdate) Validate() error {
	if strPtr := u.Name; strPtr != nil {
		str := strings.TrimSpace(*strPtr)
		if str == "" {
			return common.ErrNameCannotBeBlank
		}
		u.Name = &str
	}
	if strPtr := u.Address; strPtr != nil {
		str := strings.TrimSpace(*strPtr)
		if str == "" {
			return common.ErrAddressCannotBeBlank
		}
		u.Address = &str
	}
	return nil
}
