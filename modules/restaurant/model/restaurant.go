package restaurantmodel

import (
	"restaurantBacked/common"
)

const EntityName = "Restaurant"

type Restaurant struct {
	common.SQLModel
	Name       string             `json:"name" gorm:"column:name;"`
	Address    string             `json:"address" gorm:"column:addr;"`
	OwnerId    int                `json:"-" gorm:"column:owner_id;"`
	Logo       *common.Image      `json:"logo" gorm:"column:logo;"`
	Cover      *common.Images     `json:"cover" gorm:"column:cover;"`
	LikedCount int                `json:"liked_count" gorm:"-"`
	HasLiked   bool               `json:"has_liked" gorm:"-"`
	User       *common.SimpleUser `json:"user" gorm:"PRELOAD:false;foreignKey:OwnerId;"`
}

func (Restaurant) TableName() string { return "restaurants" }

func (data *Restaurant) Mask(dbType int) {
	data.SQLModel.Mask(dbType)
	if u := data.User; u != nil {
		u.Mask(common.DbTypeUser)
	}
}
