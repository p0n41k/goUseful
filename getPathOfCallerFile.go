package gouseful

import (
	"os"
	"runtime"
)

func GetPathOfCallerFile(osArg0 string) string {
	if len(osArg0) >= 5 && osArg0[:5] == "/tmp/" {
		if _, file, _, ok := runtime.Caller(1); ok {
			for i := len(file) - 1; i > -1; i-- {
				if file[i] == '/' {
					return file[:i+1]
				}
			}
			return file
		}
	} else {
		dir, _ := os.Getwd()
		osArg0 = dir + osArg0[1:]
		for i := len(osArg0) - 1; i > -1; i-- {
			if osArg0[i] == '/' {
				return osArg0[:i+1]
			}
		}
		return osArg0
	}
	return ""
}
