package main

import (
	"./handler"
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/shutdown", handler.ShutDown)
	http.HandleFunc("/logout", handler.Logout)
	http.HandleFunc("/restart", handler.Restart)

	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		_ = fmt.Errorf("Server Not Started. issue -> %s", err.Error())
		panic(err)
	}
}
