package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// START OMIT
	config := zap.NewDevelopmentConfig()                                // HL
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder // HL
	logger, _ := config.Build()                                         // HL
	defer logger.Sync()
	logger.Info(
		"Cannot start service.",
		zap.Int("i", 12345),
		zap.Bool("foo-enabled", false),
	)
	// END OMIT
}
