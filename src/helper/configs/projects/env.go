package projects

import (
	configs2 "github.com/faradey/madock/src/helper/configs"
	"github.com/faradey/madock/src/helper/paths"
	"github.com/faradey/madock/src/model/versions"
)

func SetEnvForProject(projectName string, defVersions versions.ToolsVersions, projectConf map[string]string) {
	generalConf := configs2.GetGeneralConfig()
	config := new(configs2.ConfigLines)
	envFile := paths.MakeDirsByPath(paths.GetExecDirPath()+"/projects/"+projectName) + "/env.txt"
	config.EnvFile = envFile
	if len(projectConf) > 0 {
		config.IsEnv = true
	}

	config.AddOrSetLine("PATH", paths.GetRunDirPath())
	config.AddOrSetLine("PLATFORM", defVersions.Platform)
	if projectConf["PLATFORM"] == "magento2" {
		Magento2(config, defVersions, generalConf, projectConf)
	} else if projectConf["PLATFORM"] == "pwa" {
		PWA(config, defVersions, generalConf, projectConf)
	} else if projectConf["PLATFORM"] == "shopify" {
		Shopify(config, defVersions, generalConf, projectConf)
	} else if projectConf["PLATFORM"] == "custom" {
		Custom(config, defVersions, generalConf, projectConf)
	}

	if !config.IsEnv {
		config.AddEmptyLine()
	}

	config.AddOrSetLine("CRON_ENABLED", configs2.GetOption("CRON_ENABLED", generalConf, projectConf))

	if !config.IsEnv {
		config.AddEmptyLine()
	}

	config.AddOrSetLine("HOSTS", defVersions.Hosts)

	if !config.IsEnv {
		config.AddEmptyLine()
	}

	config.AddOrSetLine("SSH_AUTH_TYPE", configs2.GetOption("SSH_AUTH_TYPE", generalConf, projectConf))
	config.AddOrSetLine("SSH_HOST", configs2.GetOption("SSH_HOST", generalConf, projectConf))
	config.AddOrSetLine("SSH_PORT", configs2.GetOption("SSH_PORT", generalConf, projectConf))
	config.AddOrSetLine("SSH_USERNAME", configs2.GetOption("SSH_USERNAME", generalConf, projectConf))
	config.AddOrSetLine("SSH_KEY_PATH", configs2.GetOption("SSH_KEY_PATH", generalConf, projectConf))
	config.AddOrSetLine("SSH_PASSWORD", configs2.GetOption("SSH_PASSWORD", generalConf, projectConf))
	config.AddOrSetLine("SSH_SITE_ROOT_PATH", configs2.GetOption("SSH_SITE_ROOT_PATH", generalConf, projectConf))

	if !config.IsEnv {
		config.SaveLines()
	}
}