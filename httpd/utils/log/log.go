package log

import (
	"github.com/sirupsen/logrus"
	"github.com/yangliulnn/gin-starter/configs"
	"os"
)

var Log *logrus.Logger

func Setup() {
	Log = logrus.New()
	Log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	file, err := os.OpenFile("./log/"+configs.App.Mode+".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		Log.Info("Failed to log to file, using default stderr")
	}
	Log.SetOutput(file)
	if configs.App.Mode == "debug" {
		Log.SetLevel(logrus.TraceLevel)
	}
}
