package logger

import (
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

func InitLogger() *log.Logger {
	logger := log.New()
	today := time.Now().Format("20060102")
	file, err := os.OpenFile("pts_go_"+today+".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logger.Out = file
	} else {
		logger.Info("Failed to log to file, using default stderr")
	}
	logger.SetFormatter(&log.TextFormatter{FullTimestamp: true})
	return logger
}
