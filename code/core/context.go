package core

import (
	"gorm.io/gorm"
)

type Context interface {
	GetDB() *gorm.DB
	GetBean(bean Bean) Bean
	Cleanup()

	GetControllerMap() map[string]Controller

	Run(addr ...string)
}
