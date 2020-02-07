package main

import (
	"os"

	"github.com/angelbarrera92/prometheus-multi-tenant-proxy/internal/app/prometheus-multi-tenant-proxy"
	"github.com/urfave/cli"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	app := cli.NewApp()
	app.Name = "Prometheus Multitenant Proxy"
	app.Usage = "Makes your Prometheus server multi tenant"
	app.Version = version
	app.Author = "Ángel Barrera - @angelbarrera92"
	app.Commands = []cli.Command{
		{
			Name:   "run",
			Usage:  "Runs the Prometheus multi tenant proxy",
			Action: proxy.Serve,
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "port",
					Usage: "Port to expose this prometheus proxy",
					Value: 9092,
				}, cli.StringFlag{
					Name:  "prometheus-label-proxy-endpoint",
					Usage: "Prometheus Label Proxy",
					Value: "http://localhost:9091",
				}, cli.StringFlag{
					Name:  "auth-config",
					Usage: "AuthN yaml configuration file path",
					Value: "authn.yaml",
				},
			},
		},
	}
	app.Run(os.Args)
}
