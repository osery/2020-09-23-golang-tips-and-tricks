package main

import (
	"go.uber.org/zap"
)

func main() {
	// START OMIT
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	logger.Info(
		"Cannot start service.",
		zap.Int("i", 12345),             // HL
		zap.Bool("foo-enabled", false),  // HL
	)
	// END OMIT
}
