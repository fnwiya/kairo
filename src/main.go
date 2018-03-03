package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "kairo"
	app.Usage = "warm up postgresql"
	app.Version = "0.0.1"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "database, d",
			Usage: "dbname",
		},
		cli.StringFlag{
			Name:  "table, t",
			Usage: "tablename",
		},
	}
	app.Before = func(c *cli.Context) error {
		return nil
	}
	app.Action = warmup
	app.After = func(c *cli.Context) error {
		return nil
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func warmup(c *cli.Context) {
	fmt.Println(c.String("database"))
	fmt.Println(c.String("table"))
}
