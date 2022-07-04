package main

import (
	"context"
	"github.com/cloudSlit/cloudslit/provider/internal"
	"os"

	"github.com/cloudSlit/cloudslit/provider/pkg/logger"
	"github.com/urfave/cli/v2"
)

var VERSION = "0.0.0"

func main() {
	logger.SetVersion(VERSION)
	ctx := logger.NewTagContext(context.Background(), "__main__")

	app := cli.NewApp()
	app.Name = "za-sentinel"
	app.Version = VERSION
	app.Usage = "Security, network acceleration, zero trust network architecture"
	app.Flags = commonConfig()
	app.Commands = []*cli.Command{
		newRelayCmd(ctx),
	}
	err := app.Run(os.Args)
	if err != nil {
		logger.WithContext(ctx).Errorf("%v", err)
	}
}

func newRelayCmd(ctx context.Context) *cli.Command {
	return &cli.Command{
		Name:  "relay",
		Usage: "Run relay server",
		Action: func(c *cli.Context) error {
			return internal.Run(ctx,
				internal.SetConfigFile(c.String("conf")),
				internal.SetVersion(VERSION))
		},
	}
}

func commonConfig() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:     "conf",
			Aliases:  []string{"c"},
			Usage:    "App configuration file(.json,.yaml,.toml)",
			Required: true,
		},
	}
}
