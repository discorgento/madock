package patch

import (
	"github.com/faradey/madock/src/helper/cli"
	"github.com/faradey/madock/src/helper/cli/arg_struct"
	"github.com/faradey/madock/src/helper/cli/attr"
	"github.com/faradey/madock/src/helper/configs"
	"github.com/faradey/madock/src/helper/docker"
	"github.com/faradey/madock/src/helper/logger"
	"os"
	"os/exec"
)

func Execute() {
	args := attr.Parse(new(arg_struct.ControllerGeneralPatch)).(*arg_struct.ControllerGeneralPatch)

	filePath := args.File
	patchName := args.Name
	title := args.Title
	force := args.Force

	if filePath == "" {
		logger.Fatal("The --file option is incorrect or not specified.")
	}

	projectName := configs.GetProjectName()
	projectConf := configs.GetCurrentProjectConfig()
	isForce := ""
	if force {
		isForce = "f"
	}
	service, user, workdir := cli.GetEnvForUserServiceWorkdir("php", "www-data", projectConf["workdir"])
	cmd := exec.Command("docker", "exec", "-it", "-u", user, docker.GetContainerName(projectConf, projectName, service), "php", "/var/www/scripts/php/patch-create.php", workdir, filePath, patchName, title, isForce)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		logger.Fatal(err)
	}
}
