package option

import (
	"fmt"
	"os/exec"
	"strings"
)

func ExecuteLinuxShutDownCommand() {
	app := "shutdown"

	arg0 := "now"

	cmd := exec.Command(app, arg0)

	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Print the output
	fmt.Println(string(stdout))
}

func ExecuteLinuxRestartCommand() {
	app := "reboot"

	cmd := exec.Command(app)

	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Print the output
	fmt.Println(string(stdout))
}

func ExecuteLinuxLogoutCommand() {
	whoCmd := exec.Command("who", "-u")

	stdout, err := whoCmd.Output()
	split := strings.Fields(string(stdout))
	userPid := split[len(split)-2]
	fmt.Println(userPid)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	killCmd := exec.Command("kill", userPid)
	stdout, err = killCmd.Output()

	// Print the output
	fmt.Println(string(stdout))
}
