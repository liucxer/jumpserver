package main

import (
	"github.com/liucxer/jumpserver/internal/jumpcli"
	"os"
)

func main() {
	cmd := os.Args[1:]
	cmdStr := ""
	for _, item := range cmd {
		cmdStr = cmdStr + item + " "
	}
	_ = jumpcli.Run(cmdStr)
}
