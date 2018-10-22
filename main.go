package main

import (
	"fmt"
	"log"
	"os"
	"strings"

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
		{
			Name:   "correct",
			Usage:  "make corrections based on given dict",
			Action: correct,
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

func correct(c *cli.Context) error {
	name := c.Args().Get(0)
	if name == "" {
		return fmt.Errorf("must provide name")
	}
	name = strings.ToLower(name)

	corrs := Corrections(name)
	if len(corrs) > 0 {
		alternative := corrs[0]
		fmt.Printf("Did you mean '%s' ?\n", alternative)
	}

	return nil
}
