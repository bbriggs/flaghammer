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
		// this "test then read" pattern has a race condition
		if _, err := os.Stat(flagFile); !os.IsNotExist(err) {
			// Deleting the file between os.Stat and ioutil.ReadFile
			// will produce an error that causes the program to exit
			data, err := ioutil.ReadFile(flagFile)
			if err != nil {
				fmt.Println("%s\n", err)
				os.Exit(1)
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
