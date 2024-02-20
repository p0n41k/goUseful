package gouseful

import (
	"runtime"
)

func GetPathOfCallerFile() string {
	_, path, _, _ := runtime.Caller(1)

	for i := len(path) - 1; i > -1; i-- {
		if path[i] == '/' {
			return path[:i]
		}
	}

	return path
}
