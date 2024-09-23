package middleware

import (
	"os"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		logFile, err := os.OpenFile("./logs/logfile.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}

		// Tạo một encoder cấu hình theo JSON
		encoderConfig := zap.NewDevelopmentEncoderConfig()
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		encoderConfig.TimeKey = "timestamp"
		encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

		encoder := zapcore.NewJSONEncoder(encoderConfig)
		writeSyncer := zapcore.AddSync(logFile)
		logLevel := zapcore.InfoLevel
		core := zapcore.NewCore(encoder, writeSyncer, logLevel)
		logger := zap.New(core)

		// Ghi log ví dụ
		logger.Info("Incoming Request on",
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.String("ip", c.ClientIP()),
		)

		defer logFile.Sync()
	}
}
