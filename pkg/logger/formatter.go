package logger

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
)

type formatter struct {
	serviceName string
	logrus.TextFormatter
}

func (f *formatter) getColors(entry *logrus.Entry) (int, int) {
	var (
		blueColor  = 34
		levelColor int
	)

	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = 32 // green
	case logrus.WarnLevel:
		levelColor = 33 // yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = 31 // red
	default:
		levelColor = 36 // cyan
	}

	return blueColor, levelColor
}

// Format 로그 format화
func (f *formatter) Format(entry *logrus.Entry) ([]byte, error) {
	blueColor, levelColor := f.getColors(entry)

	switch {
	case f.serviceName == "":
		if len(entry.Data) > 0 {
			return []byte(f.entryDataMessage(entry, blueColor, levelColor)), nil
		}
		return []byte(f.entryMessage(entry, blueColor, levelColor)), nil
	case len(entry.Data) > 0:
		return []byte(f.entryDataMessage(entry, blueColor, levelColor)), nil
	default:
		return []byte(f.entryMessage(entry, blueColor, levelColor)), nil
	}
}

// entryDataMessage entry Data 메세지가 있는 경우
func (f *formatter) entryDataMessage(entry *logrus.Entry, blueColor int, levelColor int) string {
	return fmt.Sprintf(
		"\x1b[%dm[%s]\x1b[0m \x1b[%dm%s\x1b[0m - %s (%s) \n",
		blueColor,
		entry.Time.Format(f.TimestampFormat),
		levelColor,
		strings.ToUpper(entry.Level.String()),
		entry.Message,
		entry.Data,
	)
}

// entryMessage entry 메시지만 처리
func (f *formatter) entryMessage(entry *logrus.Entry, blueColor int, levelColor int) string {
	return fmt.Sprintf(
		"\x1b[%dm[%s]\x1b[0m \x1b[%dm%s\x1b[0m - %s\n",
		blueColor,
		entry.Time.Format(f.TimestampFormat),
		levelColor,
		strings.ToUpper(entry.Level.String()),
		entry.Message,
	)
}
