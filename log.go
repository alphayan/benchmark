package benchmark

import (
	"github.com/rs/zerolog"
	"go.uber.org/zap"
)

func logZap() {
	logger, _ := zap.NewProduction()
	logger.Info("zap")
}
func logZapSugared() {
	logger, _ := zap.NewProduction()
	sugar := logger.Sugar()
	sugar.Infof("zap sugar")
}
func logZerolog() {
	logger := zerolog.New(zerolog.NewConsoleWriter(func(w *zerolog.ConsoleWriter) {
		w.TimeFormat = "2006-01-02 15:04:05"
	})).With().Timestamp().Caller().Logger()
	logger.Info().Msg("zerolog")
}
