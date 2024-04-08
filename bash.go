package gouseful

import (
	"crypto/sha512"
	"encoding/hex"
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

	name := getName(cmd)

	os.Remove(name)

	defer os.Remove(name)
	err := os.WriteFile(name, []byte(cmd), 0777)
	if err != nil {
		return "", err
	}

	outByte, err := exec.Command("bash", name).CombinedOutput()
	if err != nil {
		return "", err
	}

	return string(outByte), nil
}

func getName(command string) string {
	hash := sha512.New()

	hash.Write([]byte(command))
	hashed := hash.Sum(nil)
	for256Name := hex.EncodeToString(hashed)

	nameBytes := hash.Sum([]byte(for256Name))

	namestring := hex.EncodeToString(nameBytes)

	if len(namestring) > 252 {
		namestring = namestring[:252]
	}

	return namestring + ".sh"
}
