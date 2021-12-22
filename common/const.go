package common

import (
	"errors"
)

const (
	DbTypeRestaurant = 1
	DbTypeUser       = 2
	CurrentUser      = "user"
)

var (
	ErrDataNotFound         = errors.New("data not found")
	ErrNameCannotBeBlank    = errors.New("name cannot be blank")
	ErrAddressCannotBeBlank = errors.New("address cannot be blank")
)

type Requester interface {
	GetUserId() int
	GetEmail() string
	GetRole() string
}
