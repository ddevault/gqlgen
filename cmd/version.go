package cmd

import (
	"fmt"

	"git.sr.ht/~sircmpwn/gqlgen/graphql"
	"github.com/urfave/cli/v2"
)

var versionCmd = &cli.Command{
	Name:  "version",
	Usage: "print the version string",
	Action: func(ctx *cli.Context) error {
		fmt.Println(graphql.Version)
		return nil
	},
}
