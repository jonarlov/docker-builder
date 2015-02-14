package lib

import (
	"os"
	"os/exec"
	"fmt"
	"strings"
)

// ExecCommand executes the given command which uses os.Stdout and os.Stderr
// Returns any error object
func ExecCommand(command string, args ...string) {

	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	exitIfError("Error executing \"" + command + strings.Join(args, "") + "\"", err)

	err = cmd.Wait()
	exitIfError("", err)
}

// ExecOutput executes the given command and return the output
func ExecOutput(name string, args ...string) (out []byte) {

	out, err := exec.Command(name, args...).CombinedOutput()

	exitIfError(string(out), err)

	return
}

func exitIfError(message string, err error) {

	if err != nil {

		if message == "" {
			fmt.Printf("ERROR: %s\n", message)
		} else {
			fmt.Printf("ERROR: %s\n%s\n", message, err.Error())
		}

		os.Exit(1)
	}
}
