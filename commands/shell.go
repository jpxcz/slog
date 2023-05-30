package commands

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const (
	KUBECTL_NAMESPACES = `kubectl get ns  --no-headers -o custom-columns=":metadata.name"`
)

func RunShellCommand(cmdString string) []byte {
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

func RunShellCommandStdOut(cmdString string) {
	cmd := strings.Split(cmdString, " ")
	args := []string{}
	if len(cmd) > 1 {
		args = append(args, cmd[1:]...)
	}

	c := exec.Command(cmd[0], args...)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	err := c.Start()
	if err != nil {
		fmt.Printf("error starting command %s", err)
	}

	err = c.Wait()
	if err != nil {
		fmt.Printf("error at finishing command %s", err)
	}

	fmt.Printf("finished!")

}
