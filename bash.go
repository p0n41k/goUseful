package gouseful

import (
	"errors"
	"os"
	"os/exec"
)

func Bash(cmdString_or_cmdsStrArr interface{}) (string, error) {
	var cmd string
	if t, ok := cmdString_or_cmdsStrArr.([]string); ok {
		for i := 0; i < len(t); i++ {
			cmd = cmd + t[i] + "\n"
		}
	} else if t, ok := cmdString_or_cmdsStrArr.(string); ok {
		cmd = t
	} else {
		return "", errors.New("func Bash() need only string or []string")
	}
	cmd = "# !/bin/bash\n\n" + cmd

	os.Remove("tempCommands.sh")
	defer os.Remove("tempCommands.sh")
	err := os.WriteFile("tempCommands.sh", []byte(cmd), 0777)
	if err != nil {
		return "", err
	}

	outByte, err := exec.Command("bash", "tempCommands.sh").CombinedOutput()
	if err != nil {
		return "", err
	}

	return string(outByte), nil
}
