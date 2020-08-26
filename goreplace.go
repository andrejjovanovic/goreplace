package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"flag"
	"os"
	"bytes"
)

func main() {

	var configFileName string
	var manifestFileName string
	var environment string

	flag.StringVar(&manifestFileName, "m", "", "JSON manifest to parse")
	flag.StringVar(&configFileName, "c", "", "Config file to process")
	flag.StringVar(&environment, "e", "", "Environment to use")
    flag.Parse()

    if manifestFileName == "" {
        fmt.Println("Please provide JSON file by using -m option")
        return
	}
	
	if configFileName == "" {
        fmt.Println("Please provide config file by using -c option")
        return
	}

	if environment == "" {
        fmt.Println("Please provide environment by using -e option")
        return
	}

	manifestJson, err := ioutil.ReadFile(manifestFileName)
    if err != nil {
        fmt.Printf("Error reading manifest file: %s\n", err)
        return 
	}

	var result map[string]interface{}
	json.Unmarshal([]byte(manifestJson), &result)

	manifest := result[environment].(map[string]interface{})

	for key, value := range manifest {

		input, err := ioutil.ReadFile(configFileName)
			if err != nil {
					fmt.Println(err)
					os.Exit(1)
			}

		key = "{" + key + "}"

		output := bytes.ReplaceAll(input, []byte(key), []byte(value.(string)))

		if err = ioutil.WriteFile(configFileName, output, 0666); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
