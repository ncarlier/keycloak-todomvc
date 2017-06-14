package main

import (
	"os"
	"github.com/urfave/cli"
	"log"
	"todo/client"
	"todo/auth"
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
			EnvVar: "API_ENDPOINT",
		},
		cli.StringFlag{
			Name: "auth",
			Usage: "SSO endpoint",
			Value: "http://devbox/auth/",
			EnvVar: "AUTH_ENDPOINT",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:                     "list",
			ShortName:                "ls",
			Usage:                    "List my TODOs",
			Before:                startup,
			Action: func(c *cli.Context) {
				httpClient, err := client.TodoMVCClient(c.GlobalString("api"))

				if(err != nil) {
					log.Fatalf("Unable to initialize http client : %s", err)
				}

				httpClient.List()
			},
		},
		{
			Name:                     "login",
			Usage:                    "Login to remote instance",
			Before:                startup,
			Flags:			[]cli.Flag{
				cli.StringFlag{
					Name: "username, u",
					Usage: "Username",
				},
				cli.StringFlag{
					Name: "password, p",
					Usage: "Password",
				},
				cli.StringFlag{
					Name: "realm, r",
					Usage: "Realm",
					Value: "todomvc",
				},
			},
			Action: func(c *cli.Context) {
				auth.Login(c.GlobalString("auth"), c.String("username"), c.String("password"), c.String("realm"))
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