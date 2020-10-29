package ldcLog

import (
	"flag"
	"github.com/mitchellh/go-homedir"
	"log"
	"os"
	"path/filepath"
)


var (
	DefaultLogger *log.Logger
)

func init() {

	devModeFlag := flag.Bool("d", false,"development mode")

	flag.Parse()

	if *devModeFlag {
		DefaultLogger = log.New(os.Stdout, ">", log.LstdFlags )

	}else{
		var userDir, _ = homedir.Dir()

		logFilePath := filepath.FromSlash(userDir + "/liferay-docker-control.log")

		var LogFile, _ = os.OpenFile(logFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

		DefaultLogger = log.New(LogFile, ">", log.LstdFlags)
	}


}
