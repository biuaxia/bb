package support

import (
	"fmt"
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type BbLogger struct{}

func (bbl *BbLogger) Init() {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	lg := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.WarnLevel))

	zap.ReplaceGlobals(lg)

	zap.L().Info("日志加载完成")
}

func (bbl *BbLogger) InitByConfig(bbc *BbConfig) {
	writeSyncer := getLogWriterByOptions(bbc.item.LogFilename,
		bbc.item.LogMaxSize,
		bbc.item.LogMaxAge,
		bbc.item.LogMaxBackups)
	encoder := getEncoder()
	var l = new(zapcore.Level)
	err := l.UnmarshalText([]byte(bbc.item.LogLevel))
	if err != nil {
		panic(fmt.Errorf("l.UnmarshalText() failed, err: %s", err))
	}
	core := zapcore.NewCore(encoder, writeSyncer, l)

	lg := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.WarnLevel))
	zap.ReplaceGlobals(lg) // 替换zap包中全局的logger实例，后续在其他包中只需使用zap.L()调用即可

	zap.L().Info("Log 初始化完成")
}

func getLogWriter() zapcore.WriteSyncer {
	return zapcore.NewMultiWriteSyncer([]zapcore.WriteSyncer{
		zapcore.AddSync(os.Stdout),
	}...)
}

func getLogWriterByOptions(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	return zapcore.NewMultiWriteSyncer([]zapcore.WriteSyncer{
		zapcore.AddSync(&lumberjack.Logger{
			Filename:   filename,
			MaxSize:    maxSize,
			MaxBackups: maxBackup,
			MaxAge:     maxAge,
		}),
		zapcore.AddSync(os.Stdout),
	}...)
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000")
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}
