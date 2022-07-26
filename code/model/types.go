package model

import (
	"fmt"
	"strconv"
	"time"

	"biuaxia.cn/bb/code/core"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type IFormer interface {
	ConvertToModel() (m *IModel)
}

type IForm struct {
	ID        string `form:"id" json:"id"`
	CreatedAt string `form:"createdAt" json:"createdAt"`
	UpdatedAt string `form:"updatedAt" json:"updatedAt"`
}

func (iform *IForm) ConvertToModel() *IModel {
	m := &IModel{}
	if iform.ID != "" {
		id, err := strconv.Atoi(iform.ID)
		if err != nil {
			zap.L().Error("转换出错", zap.String("iform.ID", iform.ID), zap.Error(err))
			return nil
		}
		m.ID = uint(id)
	}

	m.CreatedAt = core.ParseLocalDateTime(iform.CreatedAt)
	m.UpdatedAt = core.ParseLocalDateTime(iform.UpdatedAt)

	return m
}

type IModeler interface {
	ConvertToForm() (f *IForm)
}

type IModel struct {
	ID        uint           `form:"id" gorm:"primaryKey;type:bigint(20) auto_increment;column:id;comment:唯一标识"`
	CreatedAt time.Time      `form:"createdAt" gorm:"type:datetime(3);column:created_at;comment:创建时间"`
	UpdatedAt time.Time      `form:"updatedAt" gorm:"type:datetime(3);column:updated_at;comment:更新时间"`
	DeletedAt gorm.DeletedAt `form:"deletedAt" gorm:"index;type:datetime(3);column:deleted_at;comment:删除时间"`
}

func (imodel *IModel) ConvertToForm() (f *IForm) {
	f.ID = fmt.Sprint(imodel.ID)
	loc, _ := time.LoadLocation("Local")
	// loc, _ := time.LoadLocation("Asia/Shanghai")

	f.CreatedAt = imodel.CreatedAt.In(loc).Format(core.LOCALDATETIME_FORMAT_LAYOUT)
	f.UpdatedAt = imodel.UpdatedAt.In(loc).Format(core.LOCALDATETIME_FORMAT_LAYOUT)
	return
}
