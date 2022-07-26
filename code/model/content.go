package model

import (
	"fmt"
)

type Content struct {
	IModel
	Order    int    `gorm:"column:order;type:int;comment:排序号"`
	Parent   int    `gorm:"column:parent;type:int;comment:所属文章或独立页面"`
	AuthorId uint   `gorm:"column:authorId;type:int;comment:作者唯一标识"`
	Title    string `gorm:"column:title;type:varchar(150);comment:标题"`
	Text     string `gorm:"column:text;type:longtext;comment:内容"`
	Template string `gorm:"column:template;type:varchar(32);comment:模板唯一标识"`
	Types    string `gorm:"column:type;type:varchar(16);comment:类型post文章page独立页面attachment附件"`
	Status   string `gorm:"column:status;type:varchar(16);comment:状态waiting审核中publish已发布"`
	Password string `gorm:"column:password;type:varchar(32);comment:密码"`
}

func (c *Content) ConvertToForm() (cf ContentForm) {
	imodel := c.IModel
	iform := imodel.ConvertToForm()
	cf.IForm = *iform
	cf.AuthorId = fmt.Sprint(c.AuthorId)
	return
}
