package common

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type Image struct {
	Id        int    `json:"id" gorm:"column:id;"`
	Url       string `json:"url" gorm:"column:url;"`
	Width     int    `json:"width" gorm:"column:width;"`
	Height    int    `json:"height" gorm:"column:height;"`
	CloudName string `json:"cloud_name,omitempty" gorm:"_"`
	Extension string `json:"extension,omitempty" gorm:"_"`
}

func (Image) TableName() string {
	return "image"
}

func (j *Image) Fulfill(domain string) {
	j.Url = fmt.Sprintf("%s/%s", domain, j.Url)
}

func (img *Image) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("failed to Marshal JSON value:", value))
	}
	var endImg Image
	if err := json.Unmarshal(bytes, &endImg); err != nil {
		return err
	}
	*img = endImg
	return nil
}

// Value return json value, implement driver.Value interface
func (img *Image) Value() (driver.Value, error) {
	if img == nil {
		return nil, nil
	}
	return json.Marshal(img)
}

type Images []Image

func (imgs *Images) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("failed to Marshal JSON value:", value))
	}
	var endImg []Image
	if err := json.Unmarshal(bytes, &endImg); err != nil {
		return err
	}
	*imgs = endImg
	return nil
}
func (imgs *Images) Value() (driver.Value, error) {
	if imgs == nil {
		return nil, nil
	}
	return json.Marshal(imgs)
}
