package gouseful

import (
	"os"
	"runtime"
)

func GetPathOfCallerFile(osArg0 string) string {
	if osArg0[:5] == "/tmp/" {
		if _, file, _, ok := runtime.Caller(1); ok {
			for i := len(file) - 1; i > -1; i-- {
				if file[i] == '/' {
					return file[:i+1]
				}
			}
			return file
		}
	} else {
		dir, err := os.Getwd()
		if err != nil {
			return ""
		}

		dir = dir + osArg0[1:]

		for i := len(dir) - 1; i > -1; i-- {
			if dir[i] == '/' {
				return dir[:i+1]
			}
		}
		return dir
	}
	return ""
}
