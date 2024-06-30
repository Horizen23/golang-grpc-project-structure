package main

import (
	"fmt"
	server "golang-grpc-project-structure/cmd/server/grpc"
	"log"
	"regexp"
	"strings"
)

// ANSI escape codes for color formatting
const (
	green = "\033[32m"
	reset = "\033[0m"
)

func main() {
	app, err := server.InitApp()
	if err != nil {
		log.Fatal("Failed to initialize:", err)
	}
	defer app.Disconnect()

	title := fmt.Sprintf("Welcome to the %s %s  %s", green, app.Name(), reset) 
	addr := fmt.Sprintf("Listening on %s%v (%s)%s", green, app.Addr().String(), app.Addr().Network(), reset) 
	logMessage := strings.Join([]string{
		"\n",
		"\n",
		"┌───────────────────────────────────────────────────┐",
		"│"+LogFormatCenter("", 51)+"│",
		"│"+LogFormatCenter(title, 51)+"│",
		"│"+LogFormatCenter(addr, 51)+"│",
		"│"+LogFormatCenter("", 51)+"│",
		"└───────────────────────────────────────────────────┘",
		"\n",
		"\n",
	}, "\n")

	log.Println(logMessage)
	err = app.Start()
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}


func LogFormatCenter(str string, width int) string {
	plainStr := StripANSI(str)

	contentLength := len(plainStr)

	padding := (width - contentLength) / 2

	leftPadding := ""
	if padding > 0 {
		leftPadding = fmt.Sprintf("%*s", padding, " ")
	}

	rightPadding := ""
	if padding+contentLength < width {
		rightPadding = fmt.Sprintf("%*s", width-padding-contentLength, " ")
	}

	formattedString := leftPadding + str + rightPadding

	return formattedString
}

func StripANSI(s string) string {
	re := regexp.MustCompile("\x1b\\[[0-9;]*[a-zA-Z]")
	return re.ReplaceAllString(s, "")
}