package main

import "gopkg.in/urfave/cli.v1"

var (
	serverFlags = []cli.Flag{
		cli.StringFlag{
			Name:  "port",
			Value: "3000",
			Usage: "HTTP service port",
		},
		cli.StringFlag{
			Name:  "todo-host",
			Value: "todo",
			Usage: "hostname for todo service to use",
		},
		cli.StringFlag{
			Name:  "todo-port",
			Value: "5000",
			Usage: "port for todo service to use",
		},
	}
)

type GlobalOptions struct {
	Port     string
	TodoPort string
	TodoHost string
}

func flagsFrom(flagSets ...[]cli.Flag) []cli.Flag {
	all := []cli.Flag{}
	for _, flagSet := range flagSets {
		all = append(all, flagSet...)
	}
	return all
}

func GlobalFlags() []cli.Flag {
	return flagsFrom(serverFlags)
}

func ParseOpts(c *cli.Context) GlobalOptions {
	opts := GlobalOptions{}
	opts.Port = c.GlobalString("port")
	opts.TodoHost = c.GlobalString("todo-host")
	opts.TodoPort = c.GlobalString("todo-port")
	return opts
}
