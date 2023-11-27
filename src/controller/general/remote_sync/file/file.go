package file

import (
	"fmt"
	"github.com/faradey/madock/src/configs"
	"github.com/faradey/madock/src/controller/general/remote_sync"
	"github.com/faradey/madock/src/helper/cli/attr"
	"github.com/faradey/madock/src/helper/paths"
	"github.com/jessevdk/go-flags"
	"github.com/pkg/sftp"
	"log"
	"os"
	"strings"
)

type ArgsStruct struct {
	attr.Arguments
	Path string `long:"path" short:"p" description:"Path to file on server (from site root folder)"`
}

func Execute() {
	args := getArgs()

	projectConf := configs.GetCurrentProjectConfig()
	remoteDir := projectConf["SSH_SITE_ROOT_PATH"]
	var err error
	path := strings.Trim(args.Path, "/")
	if path == "" {
		log.Fatal("")
	}
	var sc *sftp.Client
	conn := remote_sync.Connect(projectConf["SSH_AUTH_TYPE"], projectConf["SSH_KEY_PATH"], projectConf["SSH_PASSWORD"], projectConf["SSH_HOST"], projectConf["SSH_PORT"], projectConf["SSH_USERNAME"])
	fmt.Println("")
	fmt.Println("Server connection...")
	defer remote_sync.Disconnect(conn)
	sc, err = sftp.NewClient(conn)
	if err != nil {
		log.Fatal(err)
	}
	defer sc.Close()

	fmt.Println("\n" + "Synchronization is started")

	remote_sync.DownloadFile(sc, strings.TrimRight(remoteDir, "/")+"/"+path, strings.TrimRight(paths.GetRunDirPath(), "/")+"/"+path, false, false)
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
