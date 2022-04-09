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
