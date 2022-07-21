package core

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	Conf = new(AppConfig)

	Route *gin.Engine

	GormDB *gorm.DB
)
