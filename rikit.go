package main

import (
	"os"

	"github.com/mattludwigs/rikit/commands"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "Rikit"
	app.Version = "1.1.0"
	app.Usage = "API testing CLI"

	app.Commands = []cli.Command{
		commands.GET(),
	}

	app.Run(os.Args)
}
