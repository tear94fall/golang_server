package utility

import (
	"os"
	"os/exec"
	"runtime"
)

func CmdClear() {
	os_name := runtime.GOOS
	var cmd *exec.Cmd

	if os_name == "linux" {
		cmd = exec.Command("clear") //Linux
	} else if os_name == "window" {
		cmd = exec.Command("cmd", "/c", "cls") //Windows
	} else if os_name == "darwin" {
		cmd = exec.Command("clear") //Mac OS
	}

	if cmd != nil {
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
