package dao

import (
	"biuaxia.cn/bb/code/core"
	"biuaxia.cn/bb/code/model"
	"go.uber.org/zap"
	"gorm.io/gorm/clause"
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

func DeleteContent(ids ...int) {
	var contents []model.Content

	// 开始事务
	tx := core.GormDB.Begin()
	result := tx.
		Clauses(clause.Returning{}). // 返回删除行的数据
		Delete(&contents, ids)
	zap.L().Debug("DeleteContent", zap.Any("ids", ids))

	if result.Error != nil {
		zap.L().Error("DeleteContent出错", zap.Error(result.Error))
		// 回滚事务
		tx.Rollback()
		return
	}
	// 提交事务
	tx.Commit()
	zap.L().Debug("DeleteContent", zap.Any("ids", ids), zap.Any("contents", contents))
}

func CountContent() int {
	var count int64
	core.GormDB.Model(&model.Content{}).Count(&count)
	zap.L().Debug("CountContent", zap.Int64("count", count))
	return int(count)
}
