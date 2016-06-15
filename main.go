package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"gopkg.in/urfave/cli.v1"
)

var (
	startServer = cli.Command{
		Name:  "start",
		Usage: "start the todo api",
		Action: func(c *cli.Context) {
			opts := ParseOpts(c)
			api, err := NewApiService(opts.Port, opts.TodoHost, opts.TodoPort)
			if err != nil {
				log.Fatal(err)
			}
			router := api.StartServer()
			log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", opts.Port), router))
		},
	}
)

func main() {
	app := cli.NewApp()
	app.Name = "todo-api"
	app.Usage = "todo-service"
	app.Commands = []cli.Command{
		startServer,
	}
	app.Flags = GlobalFlags()
	app.Run(os.Args)
}
