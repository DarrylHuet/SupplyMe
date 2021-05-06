package main

import (
	cli "github.com/urfave/cli"

	"log"
	"os"
)

//Our Application provides the logic of our application within the framework our our Action, and it provides the main basis for our application.
func main() {
	app := &cli.App{
		Name:  "SupplyMe",
		Usage: "Decentralized Application for Logistical Services",
		Action: func(c *cli.Context) error {
			//Imput actions of our application
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
