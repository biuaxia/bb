package action

import (
	"net/http"

	"biuaxia.cn/bb/code/model"
	"biuaxia.cn/bb/code/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ContentsPostEdit(ctx *gin.Context) {
	var f model.ContentsPostEditForm
	err := ctx.ShouldBind(&f)
	if err != nil {
		zap.L().Error("绑定出错", zap.Error(err))
		return
	}

	zap.L().Debug("ContentsPostEdit", zap.Any("f", f))
	var cf model.ContentForm
	cf.Title = f.Title
	cf.Text = f.Text
	cf.CreatedAt = f.Date
	cf.Password = f.Password
	cf.Status = f.Do
	cf.Types = "post"

	service.AddContent(&cf)

	ctx.Redirect(http.StatusFound, "/admin/manage-posts")
}
