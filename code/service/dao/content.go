package dao

import (
	"biuaxia.cn/bb/code/core"
	"biuaxia.cn/bb/code/model"
	"go.uber.org/zap"
)

func InsertContent(c *model.Content) (uint, error) {
	zap.L().Debug("InsertContent", zap.Any("c", c))
	result := core.GormDB.Create(&c)
	zap.L().Debug("InsertContent", zap.Any("c", c), zap.Any("result", result))
	return c.ID, result.Error
}

func QueryAllContent() ([]model.Content, error) {
	var contents []model.Content
	result := core.GormDB.Order("id desc").Order("`order`").Find(&contents)
	zap.L().Debug("QueryAllContent", zap.Any("contents", contents), zap.Any("result", result))
	return contents, result.Error
}
