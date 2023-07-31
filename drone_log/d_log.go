package drone_log

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gookit/color"
	"log"
	"runtime"
	"strconv"
	"strings"
)

var (
	openDebug  = false
	showLineNo = false
)

func OpenDebug() {
	openDebug = true
	ShowLogLineNo(true)
}

func ShowLogLineNo(openLine bool) {
	showLineNo = openLine
}

func formatLog(msg string) string {
	if showLineNo {
		return extLogLine(msg)
	}
	return msg
}

func extLogLine(logContent string) string {
	srcFile := ""
	for i := 2; i <= 4; i++ {
		_, file, line, ok := runtime.Caller(i)

		if strings.Index(file, "d_log.go") > 0 {
			continue
		}

		if ok {
			//idx := strings.LastIndex(file, "src")
			//switch {
			//case idx >= 0:
			//	srcFile = file[idx+4:]
			//default:
			//	srcFile = file
			//}
			indexFunc := func(file string) string {
				backup := "/" + file
				lastSlashIndex := strings.LastIndex(backup, "/")
				if lastSlashIndex < 0 {
					return backup
				}
				secondLastSlashIndex := strings.LastIndex(backup[:lastSlashIndex], "/")
				if secondLastSlashIndex < 0 {
					return backup[lastSlashIndex+1:]
				}
				return backup[secondLastSlashIndex+1:]
			}
			srcFile = indexFunc(file) + ":" + strconv.Itoa(line)
		}
		break
	}
	return fmt.Sprintf("%s %s", srcFile, logContent)
}

func Debugf(format string, v ...any) {
	if !openDebug {
		return
	}
	logInfo := fmt.Sprintf(format, v...)
	log.Printf("%s %s", color.Blue.Render("debug:"), formatLog(logInfo))
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
		log.Printf("%s %s\n%s\n", color.Blue.Render("debug:"), formatLog(logInfo), dStr)
	} else {
		log.Printf("%s%s\n", color.Blue.Render("debug:"), formatLog(dStr))
	}
}

func Verbosef(format string, v ...any) {
	logInfo := fmt.Sprintf(format, v...)
	if showLineNo {
		logInfo = extLogLine(logInfo)
	}
	log.Print(logInfo)
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
		log.Printf("%s\n%s\n", formatLog(logInfo), dStr)
	} else {
		log.Printf("%s\n", formatLog(dStr))
	}
}

func Infof(format string, v ...any) {
	logInfo := fmt.Sprintf(format, v...)
	log.Printf("%s %s", color.Green.Render("info:"), formatLog(logInfo))
}

func Warnf(format string, v ...any) {
	logInfo := fmt.Sprintf(format, v...)
	log.Printf("%s %s", color.Yellow.Render("warn:"), formatLog(logInfo))
}

func Error(err error) {
	Errorf(err, "")
}

func Errorf(err error, format string, v ...any) {
	if format != "" {
		logInfo := fmt.Sprintf(format, v...)
		log.Printf("%s %s\nerror content: %v", color.Red.Render("err:"), formatLog(logInfo), err)
	} else {
		log.Printf("%s %s\n", color.Red.Render("err:"), formatLog(err.Error()))
	}
}

func Fatalf(format string, v ...any) {
	logInfo := fmt.Sprintf(format, v...)
	log.Fatalf("fatal: %s", formatLog(logInfo))
}

func Panicf(format string, v ...any) {
	logInfo := fmt.Sprintf(format, v...)
	log.Panicf("panic: %s", formatLog(logInfo))
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
		log.Printf("%s %s %v", color.Yellow.Render("warn:"), "json marshal beauty err: ", errJson)
		return "", true
	}
	dStr := str.String()
	return dStr, false
}
