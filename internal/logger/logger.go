package logger

import (
	"github.com/colevoss/temperature-blanket-backend/internal/config"
	"github.com/gin-gonic/gin"
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

func Req(c *gin.Context) *zap.SugaredLogger {
	requestId := c.GetString("requestId")

	if requestId == "" {
		return Logger
	}

	return Logger.With(zap.String("requestId", requestId))
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
