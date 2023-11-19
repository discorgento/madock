package commands

import (
	"github.com/faradey/madock/src/cli/attr"
	"github.com/faradey/madock/src/cli/fmtc"
	"github.com/faradey/madock/src/configs"
	"github.com/faradey/madock/src/docker/builder"
	"github.com/faradey/madock/src/docker/scripts"
	"github.com/faradey/madock/src/paths"
	"github.com/faradey/madock/src/ssh"
	"github.com/faradey/madock/src/versions/magento2"
)

func RemoteSyncDb() {
	projectConfig := configs.GetCurrentProjectConfig()
	conn := ssh.Connect(projectConfig["SSH_AUTH_TYPE"], projectConfig["SSH_KEY_PATH"], projectConfig["SSH_PASSWORD"], projectConfig["SSH_HOST"], projectConfig["SSH_PORT"], projectConfig["SSH_USERNAME"])
	ssh.DbDump(conn, projectConfig["SSH_SITE_ROOT_PATH"], attr.Options.Name)
}

func RemoteSyncMedia() {
	projectConfig := configs.GetCurrentProjectConfig()
	ssh.SyncMedia(projectConfig["SSH_SITE_ROOT_PATH"])
}

func RemoteSyncFile() {
	projectConfig := configs.GetCurrentProjectConfig()
	ssh.SyncFile(projectConfig["SSH_SITE_ROOT_PATH"])
}

func Proxy(flag string) {
	if !configs.IsHasNotConfig() {
		projectConfig := configs.GetCurrentProjectConfig()
		if projectConfig["PROXY_ENABLED"] == "true" {
			if flag == "prune" {
				builder.DownNginx()
			} else if flag == "stop" {
				builder.StopNginx()
			} else if flag == "restart" {
				builder.StopNginx()
				builder.UpNginx()
			} else if flag == "start" {
				builder.UpNginx()
			} else if flag == "rebuild" {
				builder.DownNginx()
				builder.UpNginx()
			}
			fmtc.SuccessLn("Done")
		} else {
			fmtc.WarningLn("Proxy service is disabled. Run 'madock service:enable proxy' to enable it")
		}
	} else {
		fmtc.WarningLn("Set up the project")
		fmtc.ToDoLn("Run madock setup")
	}
}

func Prune() {
	if !configs.IsHasNotConfig() {
		builder.Down(attr.Options.WithVolumes)
		if len(paths.GetActiveProjects()) == 0 {
			Proxy("prune")
		}
		fmtc.SuccessLn("Done")
	} else {
		fmtc.WarningLn("Set up the project")
		fmtc.ToDoLn("Run madock setup")
	}
}

func Magento(flag string) {
	builder.Magento(flag)
}

func PWA(flag string) {
	builder.PWA(flag)
}

func DebugEnable() {
	configPath := paths.GetExecDirPath() + "/projects/" + configs.GetProjectName() + "/env.txt"
	configs.SetParam(configPath, "XDEBUG_ENABLED", "true")
	Rebuild()
}

func DebugProfileEnable() {
	configPath := paths.GetExecDirPath() + "/projects/" + configs.GetProjectName() + "/env.txt"
	configs.SetParam(configPath, "XDEBUG_MODE", "profile")
	Rebuild()
}

func DebugDisable() {
	configPath := paths.GetExecDirPath() + "/projects/" + configs.GetProjectName() + "/env.txt"
	configs.SetParam(configPath, "XDEBUG_ENABLED", "false")
	Rebuild()
}

func DebugProfileDisable() {
	configPath := paths.GetExecDirPath() + "/projects/" + configs.GetProjectName() + "/env.txt"
	configs.SetParam(configPath, "XDEBUG_MODE", "debug")
	Rebuild()
}

func Info() {
	scripts.MagentoInfo()
}

func N98(flag string) {
	builder.N98(flag)
}

func Node(flag string) {
	builder.Node(flag)
}

func Logs() {
	containerName := "php"
	if len(attr.Options.Args) > 0 && attr.Options.Args[0] != "" {
		containerName = attr.Options.Args[0]
	}
	builder.Logs(containerName)
}

func IsNotDefine() {
	fmtc.ErrorLn("The command is not defined. Run 'madock help' to invoke help")
}

func Ssl() {
	builder.SslRebuild()
}

func InstallMagento() {
	toolsDefVersions := magento2.GetVersions("")
	builder.InstallMagento(configs.GetProjectName(), toolsDefVersions.Magento)
}

func MftfInit() {
	builder.MftfInit()
}

func Mftf(flag string) {
	builder.Mftf(flag)
}

func Shopify(flag string) {
	builder.Shopify(flag)
}

func ShopifyWeb(flag string) {
	builder.ShopifyWeb(flag)
}

func ShopifyWebFrontend(flag string) {
	builder.ShopifyWebFrontend(flag)
}
