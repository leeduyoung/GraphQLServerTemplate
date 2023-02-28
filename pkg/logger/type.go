package logger

import "github.com/leeduyoung/GraphQLServerTemplate/pkg"

type Config struct {
	Mode     pkg.Mode
	SlackCfg *SlackConfig
}

type SlackConfig struct {
	HookURL string
	Channel string
}
