package pkg

import (
	"os"

	"github.com/suosi-inc/go-demo/cmd/pkg/log"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// InitZapLogger Init zap logger and set to global log
func InitZapLogger() {
	logFile := viper.GetString("logger.file")
	logLevel := viper.GetString("logger.level")
	logFormat := viper.GetString("logger.format")

	// Set log level
	zapLogLevel := zapcore.DebugLevel
	switch logLevel {
	case "debug":
		zapLogLevel = zapcore.DebugLevel
	case "info":
		zapLogLevel = zapcore.InfoLevel
	case "warn":
		zapLogLevel = zapcore.WarnLevel
	case "error":
		zapLogLevel = zapcore.ErrorLevel
	case "fatal":
		zapLogLevel = zapcore.FatalLevel
	case "panic":
		zapLogLevel = zapcore.PanicLevel
	default:
		zapLogLevel = zapcore.DebugLevel
	}

	// Set log format
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006/01/02 15:04:05.999")
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	// Set log output
	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	if logFormat == "json" {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	}

	// Set log rotate
	lumberJackLogger := &lumberjack.Logger{
		Filename:   logFile,
		MaxSize:    viper.GetInt("logger.maxSize"),
		MaxBackups: viper.GetInt("logger.maxBackups"),
		MaxAge:     viper.GetInt("logger.maxAge"),
		LocalTime:  true,
		Compress:   false,
	}

	// Init zap
	fileSyncer := zapcore.AddSync(lumberJackLogger)
	consoleSyncer := zapcore.AddSync(os.Stdout)
	core := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(fileSyncer, consoleSyncer), zapLogLevel)
	zapLogger := zap.New(core, zap.AddCaller())

	// Zap to log
	zap.ReplaceGlobals(zapLogger)
	defer zapLogger.Sync()
	log.ZapToLog()

	zapLogger.Info("Init zap logger success")
}
