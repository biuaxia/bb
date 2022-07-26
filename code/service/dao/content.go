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

func UpdateContent(c *model.Content) (uint, error) {
	zap.L().Debug("UpdateContent", zap.Any("c", c))
	result := core.GormDB.Model(&c).Save(&c)
	zap.L().Debug("UpdateContent", zap.Any("c", c), zap.Any("result", result))
	return c.ID, result.Error
}

func QueryAllContent() ([]model.Content, error) {
	var contents []model.Content
	result := core.GormDB.Order("id desc").Order("`order`").Find(&contents)
	zap.L().Debug("QueryAllContent", zap.Any("contents", contents), zap.Any("result", result))
	return contents, result.Error
}

func QueryAllContentOmitText() ([]*model.Content, error) {
	var contents []*model.Content
	result := core.GormDB.Select("id",
		"created_at",
		"updated_at",
		"deleted_at",
		"IF ( LOCATE( '<!--more-->', `text` ) > 0, SUBSTR(`text`,1 ,LOCATE( '<!--more-->', `text` )-1), `text` ) AS `text`", // 匹配编辑器的分割标识
		"order",
		"parent",
		"authorId",
		"title",
		"template",
		"type",
		"status",
		"password").Order("id desc").Order("`order`").Find(&contents)
	zap.L().Debug("QueryAllContent", zap.Any("contents", contents), zap.Any("result", result))
	return contents, result.Error
}

func QueryContent(id uint) (model.Content, error) {
	var content model.Content
	result := core.GormDB.First(&content, id)
	zap.L().Debug("QueryContent", zap.Any("content", content), zap.Any("result", result))
	return content, result.Error
}
