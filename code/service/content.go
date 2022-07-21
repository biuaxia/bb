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