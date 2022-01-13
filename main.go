package main

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"text/template"
	"time"
)

var textTemplate = "template.txt"

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Please provide a pfsense config xml file as an argument")
		os.Exit(1)
	}

	file := os.Args[1]
	if _, err := os.Stat(file); errors.Is(err, os.ErrNotExist) {
		fmt.Printf("Could not find file %s\n", file)
		os.Exit(1)
	}

	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("File %s could not be opened %v", file, err)
		os.Exit(1)
	}

	pfConfig := &PFSenseConfig{}

	err = xml.Unmarshal([]byte(data), &pfConfig)
	if err != nil {
		log.Fatalf("XML decoding error %v", err)
		os.Exit(1)
	}

	t, err := template.ParseFiles(textTemplate)
	if err != nil {
		log.Fatalln(err)
	}

	now := time.Now()
	var b bytes.Buffer
	t.Execute(&b, map[string]interface{}{
		"T":      now,
		"Config": pfConfig,
	})
	fmt.Printf(b.String())

}
