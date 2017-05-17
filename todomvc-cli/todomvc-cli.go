package main

import (
	"os"
	"github.com/urfave/cli"
	"commands"
	"log"
)

func main() {
	app := cli.NewApp()
	app.Name = "todomvc"
	app.Usage = "Todo MVC"
	app.Version = "0.0.1"
	app.EnableBashCompletion = true

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "api",
			Usage: "API endpoint",
			Value: "http://devbox/api/",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:                     "list",
			ShortName:                "ls",
			Usage:                    "List my TODOs",
			Before:                startup,
			Action: func(c *cli.Context) {
				commands.List(c.GlobalString("api"))
			},
		},
	}
	app.Run(os.Args)
}

func startup(c *cli.Context) error {
	log.SetOutput(os.Stdout)
	log.SetFlags(0)
	log.SetPrefix("")

	return nil
}