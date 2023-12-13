package enable

import (
	"github.com/faradey/madock/src/controller/general/rebuild"
	"github.com/faradey/madock/src/controller/general/service"
	"github.com/faradey/madock/src/helper/cli/attr"
	configs2 "github.com/faradey/madock/src/helper/configs"
	"github.com/faradey/madock/src/helper/paths"
	"github.com/jessevdk/go-flags"
	"log"
	"os"
	"strings"
)

type ArgsStruct struct {
	attr.ArgumentsWithArgs
	Global bool `long:"global" short:"g" description:"Global"`
}

func Execute() {
	args := getArgs()

	if len(args.Args) > 0 {
		for _, name := range args.Args {
			name = strings.ToLower(name)
			if service.IsService(name) {
				serviceName := strings.ToUpper(name) + "_ENABLED"
				projectName := configs2.GetProjectName()
				envFile := ""
				if !args.Global {
					envFile = paths.MakeDirsByPath(paths.GetExecDirPath()+"/projects/"+projectName) + "/env.txt"
				} else {
					envFile = paths.MakeDirsByPath(paths.GetExecDirPath()+"/projects") + "/config.txt"
				}
				configs2.SetParam(envFile, serviceName, "true")
			}
		}
	}

	rebuild.Execute()
}

func getArgs() *ArgsStruct {
	args := new(ArgsStruct)
	if len(os.Args) > 2 {
		argsOrigin := os.Args[2:]
		var err error
		_, err = flags.ParseArgs(args, argsOrigin)

		if err != nil {
			log.Fatal(err)
		}
	}

	return args
}