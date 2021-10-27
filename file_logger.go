package nod

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const logExt = ".log"

func EnableFileLogger(dir string) (io.Closer, error) {

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return nil, err
		}
	}

	filename := filepath.Join(dir, time.Now().Format("2006-01-02-15-04-05")) + logExt
	file, err := os.Create(filename)
	if err != nil {
		return nil, err
	}
	logger := log.New(file, "", log.LstdFlags)
	fl := &fileLogger{
		file:   file,
		logger: logger,
	}

	HandleFunc(fl, LogTypes()...)

	return fl, nil
}

type fileLogger struct {
	file   io.Closer
	logger *log.Logger
}

func (fl *fileLogger) Handle(msgType MessageType, payload interface{}, topic string) {
	fl.logger.Printf(
		"%-*s %s: %v",
		maxStrLen(),
		strings.ToUpper(msgType.String()),
		strings.TrimPrefix(topic, " "),
		payload)
}

func (fl *fileLogger) Close() error {
	return fl.file.Close()
}
