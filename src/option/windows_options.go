package option

import (
	"fmt"
	"os/exec"
)

func ExecuteWindowsShutDownCommand() {
	app := "shutdown"

	arg0 := "/s"

	cmd := exec.Command(app, arg0)

	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Print the output
	fmt.Println(string(stdout))
}

func ExecuteWindowsLogoutCommand() {
	app := "shutdown"

	arg0 := "/l"

	cmd := exec.Command(app, arg0)

	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Print the output
	fmt.Println(string(stdout))
}

func ExecuteWindowsRestartCommand() {
	app := "shutdown"

	arg0 := "/r"

	cmd := exec.Command(app, arg0)

	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Print the output
	fmt.Println(string(stdout))
}
