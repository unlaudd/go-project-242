package main

import (
	"context"
	"fmt"
	"os"

	"hexlet-path-size/internal/pkg"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "hexlet-path-size",
		Usage: "print size of a file or directory",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "human",
				Aliases: []string{"H"},
				Usage:   "print human-readable sizes",
			},
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			path := c.Args().First()
			if path == "" {
				fmt.Fprintln(os.Stderr, "Error: path is required")
				os.Exit(1)
			}

			human := c.Bool("human")
			result, err := pkg.GetPathSize(path, false, human, false)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}

			fmt.Println(result)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		os.Exit(1)
	}
}
