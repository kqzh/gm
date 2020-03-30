package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

const (
	VERSION = "1.0.0"
)

func main() {
	app := cli.NewApp()
	app.Name = "EasyCrypto"
	app.Usage = "encrypt or decrypt files"
	app.Version = VERSION
	app.Commands = []cli.Command{
		encryptCommand(),
		decryptCommand(),
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
