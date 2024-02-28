package gouseful

import (
	"os/exec"
)

func GetPathOfCallerFile(osArg0 string) string {
	cmd := exec.Command("pwd")
	out, _ := cmd.CombinedOutput()

	fullpath := (string(out)[:len(out)-1] + osArg0[1:])

	for i := len(fullpath) - 1; i > -1; i-- {
		if fullpath[i] == '/' {
			return fullpath[:i]
		}
	}

	return fullpath
}
