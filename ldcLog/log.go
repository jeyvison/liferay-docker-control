package ldcLog

import (
	"github.com/mitchellh/go-homedir"
	"io"
	"log"
	"os"
	"path/filepath"
)

type LogManager struct {
	Logger       *log.Logger
	LoggerWriter io.Writer
}

var (
	DefaultLogManager LogManager
)

func init() {
	var userDir, _ = homedir.Dir()

	logFilePath := filepath.FromSlash(userDir + "/liferay-docker-control.log")

	var LogFile, _ = os.OpenFile(logFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

	var logger = log.New(LogFile, "logging: ", log.Ldate)

	stdWriter := logger.Writer()

	DefaultLogManager = LogManager{
		Logger:       logger,
		LoggerWriter: stdWriter,
	}

}
