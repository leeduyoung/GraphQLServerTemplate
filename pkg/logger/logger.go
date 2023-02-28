package logger

import (
	"fmt"
	"github.com/leeduyoung/GraphQLServerTemplate/pkg"
	"sync"
	"time"

	"github.com/lingdor/stackerror"

	"github.com/johntdyer/slackrus"
	"github.com/sirupsen/logrus"
)

var (
	log  *logrus.Logger
	once sync.Once
)

// Initialize 로그 생성
//	- Input: serviceName, cfg
//	- Output: *logrus.Logger
func Initialize(serviceName string, cfg Config) *logrus.Logger {
	if log == nil {
		once.Do(func() {
			log = logrus.New()
			log.SetFormatter(&formatter{
				serviceName,
				logrus.TextFormatter{
					FullTimestamp:          true,
					TimestampFormat:        "2006-01-02 15:04:05",
					ForceColors:            true,
					DisableLevelTruncation: true,
				},
			})

			switch cfg.Mode {
			case pkg.TestMode:
				log.SetLevel(logrus.PanicLevel)
			case pkg.DevMode:
				log.SetLevel(logrus.DebugLevel)
				addSlackHook(log, serviceName, cfg)
			case pkg.ReleaseMode:
				log.SetLevel(logrus.ErrorLevel)
				addSlackHook(log, serviceName, cfg)
			}
		})
	}
	return log
}

func addSlackHook(log *logrus.Logger, serviceName string, cfg Config) {
	if cfg.SlackCfg == nil {
		log.Warn("not found slack configuration")
		return
	}

	hookURL := cfg.SlackCfg.HookURL
	channel := cfg.SlackCfg.Channel

	acceptedLevels := []logrus.Level{
		logrus.InfoLevel,
		logrus.WarnLevel,
		logrus.ErrorLevel,
		logrus.FatalLevel,
		logrus.PanicLevel,
		logrus.DebugLevel,
	}

	log.AddHook(&slackrus.SlackrusHook{
		HookURL:        hookURL,
		AcceptedLevels: acceptedLevels,
		Channel:        channel,
		IconEmoji:      ":ghost:",
		Username:       "Ghost-bot",
		Extra: map[string]interface{}{
			"ServiceName": serviceName,
			"Timestamp":   time.Now().Local(),
		},
	})
}

// Debug 디버그
func Debug(args ...interface{}) {
	log.Debug(args...)
}

// Info ...
func Info(args ...interface{}) {
	log.Info(args...)
}

// Warn ...
func Warn(args ...interface{}) {
	log.Warn(args...)
}

// Error ...
func Error(args ...interface{}) {
	log.Error(stackerror.New(fmt.Sprint(args...)))
}

// Fatal ...
func Fatal(args ...interface{}) {
	log.Fatal(stackerror.New(fmt.Sprint(args...)))
}

// Debugf ...
func Debugf(message string, args ...interface{}) {
	log.Debugf(message, args...)
}

// Infof ...
func Infof(message string, args ...interface{}) {
	log.Infof(message, args...)
}

// Warnf ...
func Warnf(message string, args ...interface{}) {
	log.Warnf(message, args...)
}

// Errorf ...
func Errorf(message string, args ...interface{}) {
	log.Error(stackerror.New(fmt.Sprintf(message, args...)))
}

// Fatalf ...
func Fatalf(message string, args ...interface{}) {
	log.Fatalf(message, args...)
}
