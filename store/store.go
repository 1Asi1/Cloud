package store

import (
	"solution/model"
)

type Store interface {
	Create(*model.Config)
	Read(*model.Config, string, *float64) error
	Update(*model.Config) error
	Delete(float64, string, float64) error
}
