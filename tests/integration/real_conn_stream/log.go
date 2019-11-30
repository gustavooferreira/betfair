package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gustavooferreira/betfair/pkg/utils/log"
)

// Color constants
const (
	InfoColor        = "\033[0;34m%s\033[0m"
	InfoBoldColor    = "\033[1;34m%s\033[0m"
	NoticeColor      = "\033[0;36m%s\033[0m"
	NoticeBoldColor  = "\033[1;36m%s\033[0m"
	WarningColor     = "\033[0;33m%s\033[0m"
	WarningBoldColor = "\033[1;33m%s\033[0m"
	ErrorColor       = "\033[0;31m%s\033[0m"
	ErrorBoldColor   = "\033[1;31m%s\033[0m"
	DebugColor       = "\033[0;36m%s\033[0m"
	DebugBoldColor   = "\033[1;36m%s\033[0m"
)

var colorMapping = map[string]string{
	"info":         InfoColor,
	"info-bold":    InfoBoldColor,
	"notice":       NoticeColor,
	"notice-bold":  NoticeBoldColor,
	"warning":      WarningColor,
	"warning-bold": WarningBoldColor,
	"error":        ErrorColor,
	"error-bold":   ErrorBoldColor,
	"debug":        DebugColor,
	"debug-bold":   DebugBoldColor,
}

type MiniLogger struct {
	Level log.LogLevel
}

func (ml MiniLogger) Trace(msg string, fields log.Fields) {
	if ml.Level <= log.TRACE {
		timestamp := time.Now().UTC()
		line, color := generalLogging(msg, "TRACE", timestamp, fields)

		if color != "" {
			s := fmt.Sprintf(colorMapping[color], line)
			fmt.Println(s)
		} else {
			fmt.Println(line)
		}
	}
}

func (ml MiniLogger) Debug(msg string, fields log.Fields) {
	if ml.Level <= log.DEBUG {
		timestamp := time.Now().UTC()
		line, color := generalLogging(msg, "DEBUG", timestamp, fields)

		if color != "" {
			s := fmt.Sprintf(colorMapping[color], line)
			fmt.Println(s)
		} else {
			fmt.Println(line)
		}
	}
}

func (ml MiniLogger) Info(msg string, fields log.Fields) {
	if ml.Level <= log.INFO {
		timestamp := time.Now().UTC()
		line, color := generalLogging(msg, "INFO", timestamp, fields)

		if color != "" {
			s := fmt.Sprintf(colorMapping[color], line)
			fmt.Println(s)
		} else {
			fmt.Println(line)
		}
	}
}

func (ml MiniLogger) Warn(msg string, fields log.Fields) {
	if ml.Level <= log.WARN {
		timestamp := time.Now().UTC()
		line, color := generalLogging(msg, "WARN", timestamp, fields)

		if color != "" {
			s := fmt.Sprintf(colorMapping[color], line)
			fmt.Println(s)
		} else {
			fmt.Println(line)
		}
	}
}

func (ml MiniLogger) Error(msg string, fields log.Fields) {
	if ml.Level <= log.ERROR {
		timestamp := time.Now().UTC()
		line, color := generalLogging(msg, "ERROR", timestamp, fields)

		if color != "" {
			s := fmt.Sprintf(colorMapping[color], line)
			fmt.Println(s)
		} else {
			fmt.Println(line)
		}
	}
}

func generalLogging(msg string, level string, timestamp time.Time, fields log.Fields) (string, string) {
	var color = ""

	var container map[string]interface{}

	if len(fields) == 0 {
		container = map[string]interface{}{"message": msg, "level": level, "timestamp": timestamp.Format(time.RFC3339Nano)}
	} else {
		if f, ok := fields["type"]; ok {
			if f.(string) == "reader-data" {
				color = "warning-bold"
			} else if f.(string) == "writer-data" {
				color = "warning"
			}
		}
		container = map[string]interface{}{"message": msg, "level": level, "extra": fields, "timestamp": timestamp.Format(time.RFC3339Nano)}
	}

	data, err := json.Marshal(container)
	if err != nil {
		return fmt.Sprintf("{\"message\": \"%s\", \"level\":\"%s\", \"timestamp\":\"%s\"}", msg, level, timestamp.Format(time.RFC3339Nano)), color
	}

	return fmt.Sprint(string(data)), color
}
