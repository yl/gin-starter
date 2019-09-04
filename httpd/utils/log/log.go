package log

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/yangliulnn/gin-starter/configs"
)

func Setup() {
	log.SetFormatter(&log.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	file, err := os.OpenFile("./log/"+configs.App.Mode+".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Info("Failed to log to file, using default stderr")
	}
	log.SetOutput(file)

	if configs.App.Mode == "debug" {
		log.SetLevel(log.TraceLevel)
	}
}
