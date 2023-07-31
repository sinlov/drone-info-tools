package drone_log

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gookit/color"
	"log"
)

var openDebug = false

func OpenDebug() {
	openDebug = true
	ShowLogLineNo(true)
}

func ShowLogLineNo(openLine bool) {
	if openLine {
		log.SetFlags(log.Lshortfile | log.LstdFlags)
	} else {
		log.SetFlags(log.LstdFlags)
	}
}

func Debugf(format string, v ...any) {
	if !openDebug {
		return
	}
	logInfo := fmt.Sprintf(format, v...)
	log.Printf("%s %s", color.Blue.Render("debug:"), logInfo)
}

func DebugJson(d any) {
	DebugJsonf(d, "")
}

func DebugJsonf(d any, format string, v ...any) {
	if !openDebug {
		return
	}
	dStr, done := any2JsonBeauty(d)
	if done {
		return
	}
	if format != "" {
		logInfo := fmt.Sprintf(format, v...)
		log.Printf("%s %s\n", color.Blue.Render("debug:"), logInfo)
		log.Println(dStr)
	} else {
		log.Printf("%s\n%s\n", color.Blue.Render("debug:"), dStr)
	}
}

func Verbosef(format string, v ...any) {
	log.Printf(format, v...)
}

func VerboseJson(d any) {
	VerboseJsonf(d, "")
}

func VerboseJsonf(d any, format string, v ...any) {
	dStr, done := any2JsonBeauty(d)
	if done {
		return
	}
	if format != "" {
		logInfo := fmt.Sprintf(format, v...)
		log.Printf("%s\n%s\n", logInfo, dStr)
	} else {
		log.Printf("\n%s\n", dStr)
	}
}

func Infof(format string, v ...any) {
	logInfo := fmt.Sprintf(format, v...)
	log.Printf("%s %s", color.Green.Render("info:"), logInfo)
}

func Warnf(format string, v ...any) {
	logInfo := fmt.Sprintf(format, v...)
	log.Printf("%s %s", color.Yellow.Render("warn:"), logInfo)
}

func Error(err error) {
	Errorf(err, "")
}

func Errorf(err error, format string, v ...any) {
	if format != "" {
		logInfo := fmt.Sprintf(format, v...)
		log.Printf("%s %s\nerror content: %v", color.Red.Render("err:"), logInfo, err)
	} else {
		log.Printf("%s %v\n", color.Red.Render("err:"), err)
	}
}

func Fatalf(format string, v ...any) {
	logInfo := fmt.Sprintf(format, v...)
	log.Fatalf("fatal: %s", logInfo)
}

func Panicf(format string, v ...any) {
	logInfo := fmt.Sprintf(format, v...)
	log.Panicf("panic: %s", logInfo)
}

func any2JsonBeauty(d any) (string, bool) {
	data, errJson := json.Marshal(d)
	if errJson != nil {
		log.Printf("%s %s %v", color.Yellow.Render("warn:"), "json marshal err: ", errJson)
		return "", true
	}
	var str bytes.Buffer
	errBeauty := json.Indent(&str, data, "", "  ")
	if errBeauty != nil {
		log.Printf("%s %s %v", color.Yellow.Render("warn:"), "json marshal beauty err: ", errBeauty)
		return "", true
	}
	dStr := str.String()
	return dStr, false
}
