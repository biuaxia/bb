package model

import (
	"strconv"

	"go.uber.org/zap"
)

type ContentForm struct {
	IForm
	Title    string `form:"title" json:"title"`
	Text     string `form:"text" json:"text"`
	Order    string `form:"order" json:"order"`
	AuthorId string `form:"authorId" json:"authorId"`
	Template string `form:"template" json:"template"`
	Types    string `form:"types" json:"types"`
	Status   string `form:"status" json:"status"`
	Password string `form:"password" json:"password"`
	Parent   string `form:"parent" json:"parent"`
}

func (cform *ContentForm) ConvertToModel() *Content {
	c := &Content{}
	iform := cform.IForm
	zap.L().Debug("ConvertToModel", zap.Any("iform", iform))
	im := iform.ConvertToModel()
	zap.L().Debug("ConvertToModel", zap.Any("im", im))
	c.IModel = *im

	if cform.AuthorId != "" {
		authorId, err := strconv.Atoi(cform.AuthorId)
		if err != nil {
			zap.L().Error("转换出错", zap.String("cform.AuthorId", cform.AuthorId), zap.Error(err))
		}
		c.AuthorId = uint(authorId)
	}

	if cform.Parent != "" {
		parent, err := strconv.Atoi(cform.Parent)
		if err != nil {
			zap.L().Error("转换出错", zap.String("cform.AuthorId", cform.AuthorId), zap.Error(err))
		}
		c.Parent = int(parent)
	}

	c.Title = cform.Title
	c.Text = cform.Text
	c.Order = 0
	c.Template = cform.Template
	c.Types = cform.Types
	c.Status = cform.Status
	c.Password = cform.Password

	return c
}

type ContentsPostEditForm struct {
	Title        string `form:"title" json:"title"`
	Text         string `form:"text" json:"text"`
	FieldNames   string `form:"fieldNames" json:"fieldNames"`
	FieldTypes   string `form:"fieldTypes" json:"fieldTypes"`
	FieldValues  string `form:"fieldValues" json:"fieldValues"`
	Cid          string `form:"cid" json:"cid"`
	Markdown     string `form:"markdown" json:"markdown"`
	Date         string `form:"date" json:"date"`
	Tags         string `form:"tags" json:"tags"`
	Visibility   string `form:"visibility" json:"visibility"`
	Password     string `form:"password" json:"password"`
	AllowComment string `form:"allowComment" json:"allowComment"`
	AllowPing    string `form:"allowPing" json:"allowPing"`
	AllowFeed    string `form:"allowFeed" json:"allowFeed"`
	Trackback    string `form:"trackback" json:"trackback"`
	Do           string `form:"do" json:"do"`
	Timezone     string `form:"timezone" json:"timezone"`
}

type ContentsPageEditForm struct {
	Title        string `form:"title" json:"title"`
	Slug         string `form:"slug" json:"slug"`
	Text         string `form:"text" json:"text"`
	FieldNames   string `form:"fieldNames" json:"fieldNames"`
	FieldTypes   string `form:"fieldTypes" json:"fieldTypes"`
	FieldValues  string `form:"fieldValues" json:"fieldValues"`
	Cid          string `form:"cid" json:"cid"`
	Markdown     string `form:"markdown" json:"markdown"`
	Date         string `form:"date" json:"date"`
	Order        string `form:"order" json:"order"`
	Template     string `form:"template" json:"template"`
	Visibility   string `form:"visibility" json:"visibility"`
	AllowComment string `form:"allowComment" json:"allowComment"`
	AllowPing    string `form:"allowPing" json:"allowPing"`
	AllowFeed    string `form:"allowFeed" json:"allowFeed"`
	Do           string `form:"do" json:"do"`
	Timezone     string `form:"timezone" json:"timezone"`
}
