package core

import (
	"time"

	"go.uber.org/zap"
)

func ParseLocalDateTime(s string) time.Time {
	loc, _ := time.LoadLocation("Local")
	// loc, _ := time.LoadLocation("Asia/Shanghai")

	// 默认值为 golang 诞生的时间
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", "2006-01-02 15:04:05", loc)

	if s == "" {
		return t
	}

	result := LOCALDATETIME_MILLSECOND_RE.FindSubmatch([]byte(s))
	if len(result) > 0 {
		t, err := time.ParseInLocation(LOCALDATETIME_MILLSECOND_FORMAT_LAYOUT, s, loc)
		if err != nil {
			zap.L().Error("转换出错", zap.String("s", s), zap.String("layout", LOCALTIME_FORMAT_LAYOUT), zap.Error(err))
		}
		return t
	}

	result = LOCALDATETIME_MINUTE_RE.FindSubmatch([]byte(s))
	if len(result) > 0 {
		t, err := time.ParseInLocation(LOCALDATETIME_MINUTE_FORMAT_LAYOUT, s, loc)
		if err != nil {
			zap.L().Error("转换出错", zap.String("s", s), zap.String("layout", LOCALTIME_FORMAT_LAYOUT), zap.Error(err))
		}
		return t
	}

	result = LOCALDATETIME_RE.FindSubmatch([]byte(s))
	if len(result) > 0 {
		t, err := time.ParseInLocation(LOCALDATETIME_FORMAT_LAYOUT, s, loc)
		if err != nil {
			zap.L().Error("转换出错", zap.String("s", s), zap.String("layout", LOCALTIME_FORMAT_LAYOUT), zap.Error(err))
		}
		return t
	}

	result = LOCALTIME_RE.FindSubmatch([]byte(s))
	if len(result) > 0 {
		t, err := time.ParseInLocation(LOCALTIME_FORMAT_LAYOUT, s, loc)
		if err != nil {
			zap.L().Error("转换出错", zap.String("s", s), zap.String("layout", LOCALTIME_FORMAT_LAYOUT), zap.Error(err))
		}
		return t
	}

	result = LOCALDATE_RE.FindSubmatch([]byte(s))
	if len(result) > 0 {
		t, err := time.ParseInLocation(LOCALDATE_FORMAT_LAYOUT, s, loc)
		if err != nil {
			zap.L().Error("转换出错", zap.String("s", s), zap.String("layout", LOCALTIME_FORMAT_LAYOUT), zap.Error(err))
		}
		return t
	}

	t, err := time.ParseInLocation(time.RFC3339, s, loc)
	if err != nil {
		zap.L().Error("转换出错", zap.String("s", s), zap.String("layout", LOCALTIME_FORMAT_LAYOUT), zap.Error(err))
		return t
	}

	return t
}
