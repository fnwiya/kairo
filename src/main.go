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
	app.Action = func(c *cli.Context) error {
		fmt.Println("hello")
		return nil
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
