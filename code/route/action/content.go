package action

import (
	"net/http"
	"strconv"

	"biuaxia.cn/bb/code/model"
	"biuaxia.cn/bb/code/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ContentsPostEdit(ctx *gin.Context) {
	do := ctx.PostForm("do")
	if do == "" {
		do = ctx.Query("do")
	}
	zap.L().Debug("ContentsPostEdit", zap.String("do", do))
	switch do {
	case "publish": // SaveOrUpdate
		var f model.ContentsPostEditForm
		err := ctx.ShouldBind(&f)
		if err != nil {
			zap.L().Error("ContentsPostEdit-publish-绑定出错", zap.Error(err))
			ctx.Redirect(http.StatusFound, "/admin/manage-posts")
			return
		}
		zap.L().Debug("ContentsPostEdit-publish", zap.Any("f", f))
		var cf model.ContentForm
		cf.Title = f.Title
		cf.Text = f.Text
		cf.CreatedAt = f.Date
		cf.Password = f.Password
		cf.Status = f.Do
		cf.Types = "post"

		if f.Cid != "" && f.Cid != "0" {
			cf.ID = f.Cid
			// 更新
			service.RenewContent(&cf)
		} else {
			// 新增
			service.AddContent(&cf)
		}
	case "delete": // Delete, get "cid[]" form
		cids := ctx.PostFormArray("cid[]")
		zap.L().Debug("ContentsPostEdit-delete", zap.Any("cids", cids))
		if len(cids) < 1 {
			ctx.Redirect(http.StatusFound, "/admin/manage-posts")
			return
		}

		ids := make([]int, len(cids))
		for index, cid := range cids {
			ids[index], _ = strconv.Atoi(cid)
		}

		// 删除操作
		service.DelContent(ids...)
	case "mark": // Make flag, get "status" query
		zap.L().Debug("ContentsPostEdit-mark")
	}
	ctx.Redirect(http.StatusFound, "/admin/manage-posts")
}

func ContentsPageEdit(ctx *gin.Context) {
	do := ctx.PostForm("do")
	if do == "" {
		do = ctx.Query("do")
	}
	zap.L().Debug("ContentsPageEdit", zap.String("do", do))
	switch do {
	case "publish": // SaveOrUpdate
		var f model.ContentsPageEditForm
		err := ctx.ShouldBind(&f)
		if err != nil {
			zap.L().Error("ContentsPageEdit-publish-绑定出错", zap.Error(err))
			ctx.Redirect(http.StatusFound, "/admin/manage-pages")
			return
		}
		zap.L().Debug("ContentsPageEdit-publish", zap.Any("f", f))
		var cf model.ContentForm
		cf.Title = f.Title
		cf.Text = f.Text
		cf.CreatedAt = f.Date
		cf.Order = f.Order
		cf.Status = f.Do
		cf.Types = "page"

		if f.Cid != "" && f.Cid != "0" {
			cf.ID = f.Cid
			// 更新
			service.RenewContent(&cf)
		} else {
			// 新增
			service.AddContent(&cf)
		}
	case "delete": // Delete, get "cid[]" form
		cids := ctx.PostFormArray("cid[]")
		zap.L().Debug("ContentsPageEdit-delete", zap.Any("cids", cids))
		if len(cids) < 1 {
			ctx.Redirect(http.StatusFound, "/admin/manage-pages")
			return
		}

		ids := make([]int, len(cids))
		for index, cid := range cids {
			ids[index], _ = strconv.Atoi(cid)
		}

		// 删除操作
		service.DelContent(ids...)
	case "mark": // Make flag, get "status" query
		zap.L().Debug("ContentsPageEdit-mark")
	}
	ctx.Redirect(http.StatusFound, "/admin/manage-pages")
}
