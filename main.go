package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "dym"
	app.Version = "1.0.0"
	app.Commands = []cli.Command{
		{
			Name:   "variations",
			Usage:  "print variations for the given name",
			Action: vars,
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func vars(c *cli.Context) error {
	name := c.Args().Get(0)
	if name == "" {
		return fmt.Errorf("must provide name")
	}
	vs := Variations(name)
	for _, v := range vs {
		fmt.Println(v)
	}
	return nil
}
