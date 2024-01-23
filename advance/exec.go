package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	execs()
	os.Exit(1)
	env := os.Environ()
	procAttr := &os.ProcAttr{
		Env: env,
		Files: []*os.File{
			os.Stdin,
			os.Stdout,
			os.Stderr,
		},
	}
	pid, err := os.StartProcess("mvn", []string{"mvn"}, procAttr)
	if err != nil {
		fmt.Println("a error occurred")
		return
	}
	fmt.Println(pid)
}

func execs() {
	command := exec.Command("mvn")
	err := command.Run()
	if err != nil {
		fmt.Println(err)
		return
	}

}
