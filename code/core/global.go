package core

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	Conf = new(AppConfig)

	Route *gin.Engine

	GormDB *gorm.DB

	AllMenus = []Menu{
		{
			Parent: MenuItem{
				Name: "控制台",
				Href: "/admin",
			},
			Child: []MenuItem{
				{
					Name: "概要",
					Href: "/admin",
				},
				{
					Name: "个人设置",
					Href: "/admin/profile",
				},
				{
					Name: "插件",
					Href: "/admin/plugins",
				},
				{
					Name: "外观",
					Href: "/admin/themes",
				},
				{
					Name: "备份",
					Href: "/admin/backup",
				},
			},
		},
		{
			Parent: MenuItem{
				Name: "撰写",
				Href: "/admin/write-post",
			},
			Child: []MenuItem{
				{
					Name: "撰写文章",
					Href: "/admin/write-post",
				},
				{
					Name: "创建页面",
					Href: "/admin/manage-page",
				},
			},
		},
		{
			Parent: MenuItem{
				Name: "管理",
				Href: "/admin/manage-posts",
			},
			Child: []MenuItem{
				{
					Name: "文章",
					Href: "/admin/manage-posts",
				},
				{
					Name: "独立页面",
					Href: "/admin/manage-pages",
				},
				{
					Name: "评论",
					Href: "/admin/manage-comments",
				},
				{
					Name: "分类",
					Href: "/admin/manage-categories",
				},
				{
					Name: "标签",
					Href: "/admin/manage-tags",
				},
				{
					Name: "文件",
					Href: "/admin/manage-medias",
				},
				{
					Name: "用户",
					Href: "/admin/manage-users",
				},
			},
		},
		{
			Parent: MenuItem{
				Name: "设置",
				Href: "/admin/options-general",
			},
			Child: []MenuItem{
				{
					Name: "基本",
					Href: "/admin/options-general",
				},
				{
					Name: "评论",
					Href: "/admin/options-discussion",
				},
				{
					Name: "阅读",
					Href: "/admin/options-reading",
				},
				{
					Name: "永久链接",
					Href: "/admin/options-permalink",
				},
			},
		},
	}
)
