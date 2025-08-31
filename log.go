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
	fmt.Println(title + strings.Repeat("-", getTerminalWidth()-len(title)) + "\n")
}

func Fatal(str string) {
	redBg := chalk.New(chalk.BgRed).Add(chalk.FgWhite)
	redFg := chalk.New(chalk.FgRed)
	bold := chalk.New(chalk.Bold)

	errPrefix := chalk.New(chalk.Reset).Sprint() + redFg.Sprint("✘ ") + redBg.Sprint("[ERROR]") + " " + bold.Sprint(str)
	errLog := log.New(os.Stdout, errPrefix, 0)
	errLog.Fatalln()
}

func Fatalf(format string, args ...interface{}) {
	Fatal(fmt.Sprintf(format, args...))
}

func Info(str string) {
	bluBg := chalk.New(chalk.BgBlue).Add(chalk.FgWhite)
	bluFg := chalk.New(chalk.FgBlue)
	bold := chalk.New(chalk.Bold)

	errPrefix := chalk.New(chalk.Reset).Sprint() + bluFg.Sprint("▶ ") + bluBg.Sprint("[INFO]") + " " + bold.Sprint(str)
	errLog := log.New(os.Stdout, errPrefix, 0)
	errLog.Println()
}

func Infof(str string, args ...interface{}) {
	Info(fmt.Sprintf(str, args...))
}

func Warn(str string) {
	ylwBg := chalk.New(chalk.BgYellow).Add(chalk.FgWhite)
	ylwFg := chalk.New(chalk.FgYellow)
	bold := chalk.New(chalk.Bold)

	errPrefix := chalk.New(chalk.Reset).Sprint() + ylwFg.Sprint("▲ ") + ylwBg.Sprint("[WARNING]") + " " + bold.Sprint(str)
	errLog := log.New(os.Stdout, errPrefix, 0)
	errLog.Println()
}

func Warnf(str string, args ...interface{}) {
	Warn(fmt.Sprintf(str, args...))
}
