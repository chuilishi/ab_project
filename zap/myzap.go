package zap

import (
	"ab_project/global"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

// init 初始化日志库
func init() {
	encoderConfig := zap.NewProductionEncoderConfig()
	// 打印级别为大写 & 彩色
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	// 时间编码进行指定格式解析
	encoderConfig.EncodeTime = parseTime("2006-01-02 15:04:05.000")

	// 日志输出配置, 借助另外一个库 lumberjack 协助完成日志切割。
	lumberjackLogger := &lumberjack.Logger{
		Filename:   "Log.log", // -- 日志文件名
		MaxSize:    1024,      // -- 最大日志数 M为单位!!!
		MaxAge:     2,         // -- 最大存在天数
		MaxBackups: 1,         // -- 最大备份数量
		Compress:   false,     // --是否压缩
	}
	syncer := zapcore.AddSync(lumberjackLogger)

	// -- 用于开发者模式和生产模式之间的切换
	var core zapcore.Core
	if global.ENV == "debug" {
		encoder := zapcore.NewConsoleEncoder(encoderConfig)
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, syncer, zapcore.DebugLevel),
			zapcore.NewCore(encoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel),
		)
	} else {
		encoder := zapcore.NewJSONEncoder(encoderConfig)
		core = zapcore.NewCore(encoder, syncer, zapcore.InfoLevel)
	}
	lg := zap.New(core, zap.AddCaller()) // --添加函数调用信息
	zap.ReplaceGlobals(lg)               // 替换该日志为全局日志
}

// parseTime 进行时间格式处理
func parseTime(layout string) zapcore.TimeEncoder {
	return func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		type appendTimeEncoder interface {
			AppendTimeLayout(time.Time, string)
		}

		if enc, ok := enc.(appendTimeEncoder); ok {
			enc.AppendTimeLayout(t, layout)
			return
		}

		enc.AppendString(t.Format(layout))
	}
}
