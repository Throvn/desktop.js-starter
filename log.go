package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	chalk "github.com/fatih/color"
	"golang.org/x/term"
)

func getTerminalWidth() int {
	width, _, err := term.GetSize(0)
	if err != nil {
		return 50
	}
	return width
}

func PrintDivider(title string) {
	fmt.Print(title + strings.Repeat("-", getTerminalWidth()-len(title)) + "\n")
}

func Fatal(str string) {
	redFg := chalk.New(chalk.FgRed)

	errPrefix := chalk.New(chalk.Reset).Sprint() + redFg.Sprint("✘ ") + str
	errLog := log.New(os.Stdout, errPrefix, 0)
	errLog.Fatal()
}

func Fatalf(format string, args ...interface{}) {
	Fatal(fmt.Sprintf(format, args...))
}

func Info(str string) {
	bluFg := chalk.New(chalk.FgBlue)

	errPrefix := chalk.New(chalk.Reset).Sprint() + bluFg.Sprint("▶ ") + str
	errLog := log.New(os.Stdout, errPrefix, 0)
	errLog.Print()
}

func Infof(str string, args ...interface{}) {
	Info(fmt.Sprintf(str, args...))
}

func Warn(str string) {
	ylwFg := chalk.New(chalk.FgYellow)

	errPrefix := chalk.New(chalk.Reset).Sprint() + ylwFg.Sprint("▲ ") + str
	errLog := log.New(os.Stdout, errPrefix, 0)
	errLog.Print()
}

func Warnf(str string, args ...interface{}) {
	Warn(fmt.Sprintf(str, args...))
}
