package main

import (
	"github.com/integram-org/gitlab"
	"github.com/integram-org/trello"
	"github.com/integram-org/webhook"
	"github.com/kelseyhightower/envconfig"
	"github.com/requilence/integram"
)

func main() {
	var trelloCfg trello.Config
	var gitlabCfg gitlab.Config
	var webhookCfg webhook.Config

	envconfig.MustProcess("TRELLO", &trelloCfg)
	envconfig.MustProcess("GITLAB", &gitlabCfg)
	envconfig.MustProcess("WEBHOOK", &webhookCfg)

	integram.Register(
		trelloCfg,
		trelloCfg.BotConfig.Token,
	)

	integram.Register(
		gitlabCfg,
		gitlabCfg.BotConfig.Token,
	)

	integram.Register(
		webhookCfg,
		webhookCfg.BotConfig.Token,
	)

	integram.Run()
}
