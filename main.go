package main

import (
	"github.com/integram-org/gitlab"
	"github.com/integram-org/trello"
	"github.com/kelseyhightower/envconfig"
	"github.com/requilence/integram"
	"os"
)

func main() {
	integram.Debug = true

	integram.Register(
		trello.Config{
			integram.OAuthProvider{
				ID:     os.Getenv("TRELLO_APP_KEY"),
				Secret: os.Getenv("TRELLO_APP_SECRET"),
			},
		},
		os.Getenv("TRELLO_BOT_TOKEN"),
	)

	integram.Register(
		gitlab.Config{
			integram.OAuthProvider{
				ID:     os.Getenv("GITLAB_APP_ID"),
				Secret: os.Getenv("GITLAB_APP_SECRET"),
			},
		},
		os.Getenv("GITLAB_BOT_TOKEN"),
	)

	integram.Run()
}
