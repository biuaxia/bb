package service

import (
	"biuaxia.cn/bb/code/model"
	"biuaxia.cn/bb/code/service/dao"
	"go.uber.org/zap"
)

func AddContent(cf *model.ContentForm) {
	zap.L().Debug("AddContent", zap.Any("cf", cf))
	c := cf.ConvertToModel()
	id, err := dao.InsertContent(c)
	if err != nil {
		zap.L().Error("新增出错", zap.Error(err))
		return
	}
	zap.L().Debug("AddContent", zap.Uint("id", id))
}

func RenewContent(cf *model.ContentForm) {
	zap.L().Debug("RenewContent", zap.Any("cf", cf))
	c := cf.ConvertToModel()
	id, err := dao.UpdateContent(c)
	if err != nil {
		zap.L().Error("更新出错", zap.Error(err))
		return
	}
	zap.L().Debug("RenewContent", zap.Uint("id", id))
}

func GetAllContent() []model.Content {
	contents, err := dao.QueryAllContent()
	if err != nil {
		zap.L().Error("查询出错", zap.Error(err))
		return nil
	}
	zap.L().Debug("GetAllContent", zap.Any("contents", contents))
	return contents
}

func GetAllContentOmitText() []*model.Content {
	contents, err := dao.QueryAllContentOmitText()
	if err != nil {
		zap.L().Error("查询出错", zap.Error(err))
		return nil
	}

	zap.L().Debug("GetAllContent", zap.Any("contents", contents))
	return contents
}

func GetContent(id uint) model.Content {
	content, err := dao.QueryContent(id)
	if err != nil {
		zap.L().Error("查询出错", zap.Error(err))
		return model.Content{}
	}
	zap.L().Debug("GetContent", zap.Any("content", content))
	return content
}

func ViewContentByArchives(id uint) model.Content {
	return GetContent(id)
}
