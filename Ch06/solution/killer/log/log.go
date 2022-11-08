package log

import (
	"go.uber.org/zap"
)

// New returns a new logger.
// New might panic in some edge cases.
func New(name string) *zap.SugaredLogger {
	cfg := zap.NewProductionConfig()
	cfg.InitialFields = map[string]any{
		"name": name,
	}

	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}

	return logger.Sugar()
}
