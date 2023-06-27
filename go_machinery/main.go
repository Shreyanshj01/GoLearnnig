package main

import (
	"os"

	"github.com/RichardKnop/machinery/v1"
	"github.com/urfave/cli"

	"go-machinery/server"
	"go-machinery/utils"
	"go-machinery/worker"
)

var (
	app        *cli.App
	taskserver *machinery.Server
)

func init() {
	app = cli.NewApp()
	taskserver = utils.GetMachineryServer()
}

func main() {
	app.Commands = []cli.Command{
		{
			Name:  "server",
			Usage: "Run the server that takes task input",
			Action: func(c *cli.Context) {
				server.StartServer(taskserver)
			},
		},
		{
			Name:  "worker",
			Usage: "Run the worker that will consume tasks",
			Action: func(c *cli.Context) {
				worker.StartWorker(taskserver)
			},
		},
	}
	app.Run(os.Args)
}
