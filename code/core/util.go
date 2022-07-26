package core

import (
	"time"

	"go.uber.org/zap"
)

func ParseLocalDateTime(s string) (t time.Time) {
	if s == "" {
		return t
	}
	loc, _ := time.LoadLocation("Local")
	// loc, _ := time.LoadLocation("Asia/Shanghai")

	result := LOCALDATETIME_MILLSECOND_RE.FindSubmatch([]byte(s))
	if len(result) > 0 {
		t, err := time.ParseInLocation(LOCALDATETIME_MILLSECOND_FORMAT_LAYOUT, s, loc)
		if err != nil {
			zap.L().Error("转换出错", zap.String("s", s), zap.Error(err))
		}
		return t
	}

	result = LOCALDATETIME_MINUTE_RE.FindSubmatch([]byte(s))
	if len(result) > 0 {
		t, err := time.ParseInLocation(LOCALDATETIME_MINUTE_FORMAT_LAYOUT, s, loc)
		if err != nil {
			zap.L().Error("转换出错", zap.String("s", s), zap.Error(err))
		}
		return t
	}

	result = LOCALDATETIME_RE.FindSubmatch([]byte(s))
	if len(result) > 0 {
		t, err := time.ParseInLocation(LOCALDATETIME_FORMAT_LAYOUT, s, loc)
		if err != nil {
			zap.L().Error("转换出错", zap.String("s", s), zap.Error(err))
		}
		return t
	}

	return t
}
