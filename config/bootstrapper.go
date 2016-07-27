package config

import "github.com/urfave/cli"

const (
	Name   = "merged-prs"
	Author = "@promiseofcake"
)

type Bootstrapper struct {
	Cfg *Config
	*cli.App
}

func Bootstrap() *Bootstrapper {
	return &Bootstrapper{
		App: newApp(),
		Cfg: NewDefault(),
	}
}

func newApp() *cli.App {
	a := cli.NewApp()
	a.Name = Name
	a.Author = Author
	return a
}
