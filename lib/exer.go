package lib

import (
	"os"
	"os/exec"
)

// ExecCommand executes the given command which uses os.Stdout and os.Stderr
// Returns any error object
func ExecCommand(name string, args ...string) (err error) {

	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Start()

	if err != nil {
		return
	}

	err = cmd.Wait()

	return
}

// ExecOutput executes the given command and return the output
func ExecOutput(name string, args ...string) (out []byte, err error) {

	out, err = exec.Command(name, args...).CombinedOutput()

	return
}
