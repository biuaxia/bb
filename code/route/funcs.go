package route

import (
	"html/template"
	"time"
)

var (
	adminFuncMap = template.FuncMap{
		"parseDateTime": parseDateTime,
	}
	openapiFuncMap = template.FuncMap{
		"copy": func() string {
			return time.Now().Format("2006")
		},
	}
)

const (
	TimeFormat = "2006-01-02 15:04:05"
)

func parseDateTime(t time.Time) string {
	return t.Format(TimeFormat)
}
