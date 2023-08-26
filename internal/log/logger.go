package log

import (
	"context"

	"github.com/colevoss/temperature-blanket-backend/internal/config"
	"go.uber.org/zap"
)

var Logger *zap.SugaredLogger

func InitLogger(cfg *config.Config) {
	loggerConfig := getLoggerConfig(cfg)

	logger := zap.Must(loggerConfig.Build())

	Logger = logger.Sugar()
}

func CloseLogger() {
	if Logger != nil {
		_ = Logger.Sync()
	}
}

func Raw() *zap.SugaredLogger {
	return Logger
}

func C(ctx context.Context) *zap.SugaredLogger {
	requestId := ctx.Value("requestId")

	if requestId == nil {
		return Logger
	}

	return Logger.With(zap.String("requestId", requestId.(string)))
}

func getLoggerConfig(cfg *config.Config) zap.Config {
	var loggerConfig zap.Config

	if cfg.IsProd() {
		loggerConfig = zap.NewProductionConfig()
	} else {
		loggerConfig = zap.NewDevelopmentConfig()
	}

	loggerConfig.EncoderConfig.MessageKey = "message"

	return loggerConfig
}
