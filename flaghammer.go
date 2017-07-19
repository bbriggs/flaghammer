package main

import (
	"io/ioutil"
	"os"
	"bytes"
	"fmt"
)

func main() {
	args := os.Args[1:]
	flagFile := args[0]
	flagString := []byte(args[1])
	for {
		if _, err := os.Stat(flagFile); !os.IsNotExist(err) {
			data, err := ioutil.ReadFile(flagFile)
			if err != nil {
				fmt.Println("%s\n", err)
				os.Exit(1)
			}
			if !bytes.Equal(data, flagString) {
				ioutil.WriteFile(flagFile, flagString, 0644)
			}
		} else {
			ioutil.WriteFile(flagFile, flagString, 0644)
		}
	}
}
