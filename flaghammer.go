package main

import (
	"bytes"
	"fmt"
	"github.com/bbriggs/go-utils"
	"github.com/urfave/cli"
	"io/ioutil"
	"os"
	"time"
)

func main() {
	app := cli.NewApp()
	app.Version = "0.1.1"
	app.Usage = "For when you absolutely, positively need to write a file over and over again."
	app.HelpName = "Flaghammer"
	app.UsageText = "flaghammer filename string-to-write"
	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Bren \"fraq\" Briggs",
			Email: "code@fraq.io",
		},
	}
	app.Action = func(c *cli.Context) error {
		if c.NArg() < 2 {
			fmt.Errorf("Error: Not enough arguments.")
		}
		flagfile := c.Args()[0]
		flagstring := []byte(c.Args()[1])
		hammer(flagfile, flagstring)
		utils.WaitForCtrlC()
		return nil
	}
	app.Run(os.Args)
}

func hammer(flagFile string, flagString []byte) error {
	for {
		// this "test then read" pattern has a race condition
		if _, err := os.Stat(flagFile); !os.IsNotExist(err) {
			// Deleting the file between os.Stat and ioutil.ReadFile
			// will produce an error that causes the program to exit
			data, err := ioutil.ReadFile(flagFile)
			if err != nil {
				return err
			}
			// Only write when file exists and doesn't match flagString
			if !bytes.Equal(data, flagString) {
				ioutil.WriteFile(flagFile, flagString, 0644)
			}
		} else {
			ioutil.WriteFile(flagFile, flagString, 0644)
		}
	}
}
