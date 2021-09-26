package main

import (
	"fmt"
	"github.com/faradey/madock/src/cli/commands"
	"github.com/faradey/madock/src/cli/helper"
	"github.com/faradey/madock/src/configs"
	"github.com/faradey/madock/src/paths"
	"os"
	"strings"
)

func main() {
	if len(os.Args) > 1 {
		command := strings.ToLower(os.Args[1])
		switch command {
		case "setup":
			commands.Setup()
		case "start":
		case "stop":
		case "restart":
		case "refresh":
		case "magento":
		case "composer":
		case "dbimport":
		case "dbexport":
		case "help":
			helper.Help()
		default:
			commands.IsNotDefine()
		}
		fmt.Println(command)
		fmt.Println(paths.GetExecDirName())
		fmt.Println(paths.GetExecDirPath())
		fmt.Println(paths.GetRunDirPath())
		fmt.Println(paths.GetRunDirName())
		fmt.Println(configs.GetGeneralConfig())
	} else {
		helper.Help()
	}
}
