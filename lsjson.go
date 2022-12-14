package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// Read the entries in the current directory
	entries, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Check if the directory is empty
	if len(entries) == 0 {
		fmt.Println("[]")
		return
	}

	// Marshal the entries into a JSON array
	json, err := json.Marshal(entries)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the JSON array
	os.Stdout.Write(json)
}
