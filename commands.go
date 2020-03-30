package main

import (
	"github.com/urfave/cli"
)

func encryptCommand() cli.Command {
	return cli.Command{
		Name: "encrypt",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "file, f",
				Usage: "e.g: --file a.txt",
			},
			cli.StringFlag{
				Name:  "key, k",
				Usage: "e.g: --key xxx",
				Value: "0123456789abcdeffedcba9876543210",
			},
			cli.StringFlag{
				Name:  "InitVector, iv",
				Value: "0123456789abcdeffedcba9876543210",
				Usage: "e.g: --InitVector xxx",
			},
		},
		Action: func(ctx *cli.Context) {
			iv := ctx.String("iv")
			key := ctx.String("key")
			file := ctx.String("file")
			Encrypt(file, key, iv)
		},
	}
}

func decryptCommand() cli.Command {
	return cli.Command{
		Name: "decrypt",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "file, f",
				Usage: "e.g: --file a.txt",
			},
			cli.StringFlag{
				Name:  "key, k",
				Usage: "e.g: --key xxx",
				Value: "0123456789abcdeffedcba9876543210",
			},
		},
		Action: func(ctx *cli.Context) {
			key := ctx.String("key")
			file := ctx.String("file")
			Decrypt(file, key)
		},
	}
}
