package logger

import (
	"fmt"
	"os"

	"github.com/calmonr/cicd/internal/cli"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func New() (*zap.Logger, error) {
	level, err := zap.ParseAtomicLevel(viper.GetString(cli.LogLevelFlag))
	if err != nil {
		return nil, fmt.Errorf("could not parse level: %w", err)
	}

	logger := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		zapcore.Lock(os.Stdout),
		level,
	))

	return logger, nil
}
