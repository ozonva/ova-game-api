package logs

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"time"
)

var FileLogger *os.File

func InitLogger() {
	file, err := os.OpenFile("logs/app-single-logs.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		consoleWriter := zerolog.ConsoleWriter{Out: file, TimeFormat: time.RFC3339}
		multi := zerolog.MultiLevelWriter(consoleWriter, os.Stdout)
		log.Logger = zerolog.New(multi).With().Timestamp().Logger()
	}
	FileLogger = file
}
