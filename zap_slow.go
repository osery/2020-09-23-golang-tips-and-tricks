package main

import (
	"go.uber.org/zap"
)

func main() {
	// START OMIT
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar() // HL
	sugar.Infow(
		"Cannot start service.",
		"i", 12345, // HL
		"foo-enabled", false, // HL
	)
	// END OMIT
}
