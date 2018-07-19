package main

import (
	"log"
	"os"

	"github.com/pipizhang/pubsub/router"
	"github.com/urfave/cli"
)

var appName = `
  ╔═╗┬ ┬┌┐ ╔═╗┬ ┬┌┐
  ╠═╝│ │├┴┐╚═╗│ │├┴┐
  ╩  └─┘└─┘╚═╝└─┘└─┘
`
var defaultConfFile = "conf/conf.toml"
var defaultIPDBFile = "conf/ipfilter-GeoLite2-Country.mmdb.gz"

func main() {

	app := cli.NewApp()
	app.Name = appName
	app.Usage = ""
	app.Version = "1.0"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Value: defaultConfFile,
			Usage: "Load configuration file from `FILE`",
		},
		cli.StringFlag{
			Name:  "ipdb, i",
			Value: defaultIPDBFile,
			Usage: "Load IP DB from `FILE`",
		},
	}

	app.Action = func(c *cli.Context) {
		router.Start(c.String("config"), c.String("ipdb"))
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
