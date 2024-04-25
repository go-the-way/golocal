package golocal

import (
	"regexp"
	"runtime"
	"strings"
)

var re = regexp.MustCompile(`^goroutine (\d+) \[\w+\]:$`)

func GoId() string {
	var buf = make([]byte, 100)
	n := runtime.Stack(buf, false)
	// goroutine 35 [running]:
	// main.main.func1.main.func1.1.2()
	//        /Volumes/Extra/Desktop/duckcp/app/cmd/main
	str := string(buf[:n])
	return re.ReplaceAllString(strings.Split(str, "\n")[0], "$1")
}
