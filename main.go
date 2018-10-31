package main

import (
	"fmt"
	"io/ioutil"
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
		{
			Name:   "server",
			Usage:  "start corrections http server",
			Flags:  []cli.Flag{cli.StringFlag{Name: "port", Value: "8080", Usage: "http port"}},
			Action: DymServer,
		},
		{
			Name:   "check",
			Usage:  "check for corrections in a list of names",
			Flags:  []cli.Flag{cli.StringFlag{Name: "file", Value: "", Usage: "file of names"}},
			Action: check,
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

func check(c *cli.Context) error {
	file := c.String("file")
	if len(file) <= 0 {
		return fmt.Errorf("missing --file argument")
	}
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	lines := strings.Split(string(data), "\n")
	for _, l := range lines {
		s := strings.TrimSpace(l)
		s = strings.ToLower(s)
		if len(s) == 0 {
			continue
		}
		corrs := Corrections(s)
		if len(corrs) > 0 {
			fmt.Printf("would suggest '%s' for '%s'\n", corrs[0], s)
		}
	}
	return nil
}
