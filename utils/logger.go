package utils

import (
	"strings"

	"github.com/fatih/color"
)

type logger struct{}

// Infoln logs an info message
func (l *logger) Infoln(s ...string) {
	color.New(color.FgCyan).Println("[info]", strings.Join(s, " "))
}

// Warnln logs an warning message
func (l *logger) Warnln(s ...string) {
	color.New(color.FgYellow).Println("[warn]", strings.Join(s, " "))
}

// Errorln logs an error message
func (l *logger) Errorln(s ...string) {
	color.New(color.FgRed).Println("[err]", strings.Join(s, " "))
}

// Successln logs a success message
func (l *logger) Successln(s ...string) {
	color.New(color.FgGreen).Println("[success]", strings.Join(s, " "))
}

// LoggerService is the default logger instance
var LoggerService = &logger{}
