package commands

import (
	"fmt"
	"os/exec"
	"strings"
)

func RunShellCommand(cmdString string) []byte {
	// var stdout bytes.Buffer
	cmd := strings.Split(cmdString, " ")
	args := []string{}
	if len(cmd) > 1 {
		args = append(args, cmd[1:]...)
	}

	c := exec.Command(cmd[0], args...)
	output, err := c.CombinedOutput()
	if err != nil {
		fmt.Printf("Error executing command '%s', %s", cmd, err)
	}

	return output
}
