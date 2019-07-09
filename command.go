package command

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os/exec"
	"strings"
)

func containsArg(slice []string, arg string) (bool, int) {
	for index := range slice {
		if strings.Contains(slice[index], arg) {
			return true, index
		}
	}
	return false, 0
}

func Run(cmd string) (string, error) {
	var stdout []byte
	slice := strings.Split(cmd, " ")
	containsPipe, index := containsArg(slice, "|")

	// Works only for 1 pipe at the moment
	if containsPipe {
		// split slice into two parts. One before the pipe and one after

		beforePipe := slice[0 : index-1]
		afterPipe := slice[index+1:]

		// See https://golang.org/pkg/os/exec/#Cmd.StdinPipe
		c1 := exec.Command(beforePipe[0], beforePipe[1:]...)
		c2 := exec.Command(afterPipe[0], afterPipe[1:]...)
		r, w := io.Pipe()
		c1.Stdout = w
		c2.Stdin = r
		var b2 bytes.Buffer
		c2.Stdout = &b2
		c1.Start()
		c2.Start()
		c1.Wait()
		w.Close()
		c2.Wait()
		str := ""
		str = b2.String()
		if str == "" {
			err := errors.New("grep didn't return any result")
			return str, err
		}
		return str, nil
	} else {
		stdout, err := exec.Command(slice[0], slice[1:]...).Output()
		if err != nil {
			println(err.Error())
			return string(stdout), err
		}
	}
	return string(stdout), nil
}

func debugRun(cmd string) {
	slice := strings.Split(cmd, " ")
	str := ""
	for index, element := range slice {
		fmt.Printf("%4d : %v\n", index, element)
		str = str + fmt.Sprintf("%s ", element)
	}
	fmt.Printf("The command is: %s\n", str)
}
