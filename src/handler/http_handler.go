package handler

import (
	"../option"
	"fmt"
	"net/http"
	"runtime"
)

func ShutDown(w http.ResponseWriter, req *http.Request) {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("ExecuteMacOsShutDownCommand.")
		option.ExecuteMacOsShutDownCommand()
	case "linux":
		fmt.Println("ExecuteLinuxShutDownCommand.")
		option.ExecuteLinuxShutDownCommand()
	case "windows":
		fmt.Println("ExecuteWindowsShutDownCommand.")
		option.ExecuteWindowsShutDownCommand()
	default:
		fmt.Printf("there is no option for this os (%s).\n", os)
	}
}
func Restart(w http.ResponseWriter, req *http.Request) {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("ExecuteMacOsRestartCommand.")
		option.ExecuteMacOsRestartCommand()
	case "linux":
		fmt.Println("ExecuteLinuxRestartCommand.")
		option.ExecuteLinuxRestartCommand()
	case "windows":
		fmt.Println("ExecuteWindowsRestartCommand.")
		option.ExecuteWindowsRestartCommand()
	default:
		fmt.Printf("there is no option for this os (%s)", os)
	}
}
func Logout(w http.ResponseWriter, req *http.Request) {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("ExecuteMacOsLogoutCommand.")
		option.ExecuteMacOsLogoutCommand()
	case "linux":
		fmt.Println("ExecuteLinuxLogoutCommand.")
		option.ExecuteLinuxLogoutCommand()
	case "windows":
		fmt.Println("ExecuteWindowsLogoutCommand.")
		option.ExecuteWindowsLogoutCommand()
	default:
		fmt.Printf("there is no option for this os (%s)", os)
	}
}
