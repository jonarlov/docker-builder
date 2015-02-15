package lib

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
)

type ExecResult struct {

	stdout []byte
	stderr []byte
	err error
	errorMessage string
}

// Return stdout and stderr combined as string
func (e *ExecResult) Out() (out string) {
	e.ErrorHandling()
	return string(e.stdout) + string(e.stderr)
}

// Do error handling
func (e *ExecResult) ErrorHandling() {

	if e.err != nil {

		if e.errorMessage == "" {
			fmt.Errorf("%s\n", e.err)
		} else {
			fmt.Errorf("%s\n%s", e.errorMessage, e.err)
		}
	}
}

// ExecCommand executes the given command which uses os.Stdout and os.Stderr
// Returns whatever was printed on stdout and stderr
func ExecCommand(command string, args []string) (result ExecResult) {

	fmt.Printf("%s %s\n", command, strings.Join(args, " "))

	cmd := exec.Command(command, args...)
	stdout, err := cmd.StdoutPipe()
	if err != nil { return ExecResult{err: err} }
	stderr, err := cmd.StderrPipe();
	if err != nil { return ExecResult{err: err} }

	err = cmd.Start()
	if err != nil { return ExecResult{errorMessage: "Error executing \""+command+strings.Join(args, "")+"\"", err: err} }

	// read all output and print it
	outStd, err := ioutil.ReadAll(stdout)
	if err != nil { return ExecResult{err: err} }
	errStd, err := ioutil.ReadAll(stderr)
	if err != nil { return ExecResult{err: err} }

	// print outStd and errStd on stdOut
	fmt.Println(string(outStd))
	fmt.Println(string(errStd))

	err = cmd.Wait()

	return ExecResult{stdout: outStd, stderr: errStd, err: err}
}

// ExecOutput executes the given command and return the output
func ExecOutput(command string, args []string) (result ExecResult) {

	fmt.Printf("%s %s\n", command, strings.Join(args, " "))

	out, err := exec.Command(command, args...).CombinedOutput()
	if err != nil { return ExecResult{stdout: out, err: err} }

	return ExecResult{stdout: out}
}
