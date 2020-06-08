package main

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type File struct {
	Input     string
	Delimeter string
	Output    string
}

type Unixfiles struct {
	Files []File
}

func openFile(filename string) (*os.File, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func createCsv(file string) (*os.File, error) {
	f, err := os.Create(file)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func processData(metadata interface{}, b []byte) {
	var csvd = ","
	json.Unmarshal(b, &metadata)
	m := metadata.(map[string]interface{})

	for _, value := range m {
		for _, files := range value.([]interface{}) {
			record := files.(map[string]interface{})
			inputfile, err := openFile(record["Input"].(string))
			if err != nil {
				log.Fatal(err)
				os.Exit(1)
			}
			defer inputfile.Close()

			d := record["Delimeter"].(string)
			lineoutput := bufio.NewScanner(inputfile)
			outputfile, err := createCsv(record["Output"].(string))
			for lineoutput.Scan() {
				if strings.Contains(lineoutput.Text(), d) {
					newstr := strings.Replace(lineoutput.Text(), d, csvd, -1)
					_, err := outputfile.WriteString(newstr + "\n")
					if err != nil {
						log.Fatal(err)
						os.Exit(1)
					}
			                defer outputfile.Close()
				}
			}
		}
	}
}

func main() {

	// Open json file
	jsonfile, err := openFile("unixmeta.json")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer jsonfile.Close()

	// Read json file into a byte array
	b, err := ioutil.ReadAll(jsonfile)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	processData(Unixfiles{}, b)
}
