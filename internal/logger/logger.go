package logger

import (
	"log"

	"go.uber.org/zap"
)

func InitLogger() {
	l, err := zap.NewProduction()
	if err != nil {
		log.Panic("Failed to initialize zap logger")
	}
	zap.ReplaceGlobals(l)
	defer l.Sync()
}

func INFO(msg string) {
	zap.L().Info(msg)
}

func ERROR(msg, err string) {
	zap.L().Error(msg, zap.Any("error", err))
}
