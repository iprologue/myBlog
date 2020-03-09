package logger

import (
	"github.com/gin-gonic/gin"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"log"
	"time"
)

func LoggerToFile() gin.HandlerFunc {

	fileName := getLogFileName()
	filePath := getLogFilePath()
	file, err := openLogFile(fileName, filePath)
	if err != nil {
		log.Fatalln(err)
	}

	logger := logrus.New()

	logger.Out = file
	logger.SetLevel(logrus.DebugLevel)

	logs, err := rotatelogs.New(
		filePath+fileName,
		rotatelogs.WithLinkName(filePath),
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	if err != nil {
		log.Fatalf("rotatelogs err: %v", err)
	}

	writerMaps := lfshook.WriterMap{
		logrus.InfoLevel:  logs,
		logrus.FatalLevel: logs,
		logrus.DebugLevel: logs,
		logrus.WarnLevel:  logs,
		logrus.ErrorLevel: logs,
		logrus.PanicLevel: logs,
	}

	lfsHook := lfshook.NewHook(writerMaps, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	logger.AddHook(lfsHook)

	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next()

		endTime := time.Now()

		latencyTime := endTime.Sub(startTime)

		reqMethod := c.Request.Method

		reqUri := c.Request.RequestURI

		statusCode := c.Writer.Status()

		clientIP := c.ClientIP()

		logger.WithFields(logrus.Fields{
			"status_code":  statusCode,
			"latency_time": latencyTime,
			"client_ip":    clientIP,
			"req_method":   reqMethod,
			"req_uri":      reqUri,
		}).Info()
	}
}
