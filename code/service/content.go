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

func GetAllContent(typeStr string) []model.ContentForm {
	contents, err := dao.QueryAllContent(typeStr)
	if err != nil {
		zap.L().Error("查询出错", zap.Error(err))
		return nil
	}

	var forms []model.ContentForm
	for _, content := range contents {
		forms = append(forms, content.ConvertToForm())
	}
	zap.L().Debug("GetAllContent", zap.Any("forms", forms))

	return forms
}

func GetAllContentOmitText(typeStr string) []*model.ContentForm {
	contents, err := dao.QueryAllContentOmitText(typeStr)
	if err != nil {
		zap.L().Error("查询出错", zap.Error(err))
		return nil
	}

	var forms []*model.ContentForm
	for _, content := range contents {
		form := content.ConvertToForm()
		forms = append(forms, &form)
	}
	zap.L().Debug("GetAllContent", zap.Any("forms", forms))

	return forms
}

func GetContent(id uint) model.ContentForm {
	content, err := dao.QueryContent(id)
	if err != nil {
		zap.L().Error("查询出错", zap.Error(err))
		return model.ContentForm{}
	}
	zap.L().Debug("GetContent", zap.Any("content", content))
	return content.ConvertToForm()
}

func ViewContentByArchives(id uint) model.ContentForm {
	return GetContent(id)
}

func DelContent(ids ...int) {
	dao.DeleteContent(ids...)
}

func AllContentTotal(typeStr string) int {
	return dao.CountContent(typeStr)
}
