package main 

import (
    "fmt"
    "log"
    "os"
    "bufio"
    "io/ioutil"
    "encoding/json"
)


func main() {

    // var csvd string = ","

    type File struct {
        Input string `json:"inputfile"`  
	Delimeter string `json:"delimeter"` 
	Output string `json:"outputfile"` 
    }

    type Unixfiles struct  {
	Files []File 
    }

    // Open json file
    jsonfile, err := openFile("unixmeta.json")
    if err != nil {
	log.Fatal(err)
        os.Exit(1) 
    }

    // Read json file into a byte array
    b, err := ioutil.ReadAll(jsonfile)
    if err != nil {
	log.Fatal(err)
        os.Exit(1) 
    }

    var data Unixfiles 
    err := json.Unmarshal(b, &data)
    if (err != nil) {
   		return
	}
    fmt.Printf("%T\n", data.Files)
    processData(data)
    jsonfile.Close()
}

func openFile(filename string) (*os.File, error) {
        f, err := os.Open(filename)
        if err != nil {
		return nil, err 
        }
        return f, nil
}

func processData(metadata Unixfiles) {

    for _,value := range metadata.Files {
      file, err := openFile(value.Input)
      if err != nil {
	log.Fatal(err)
        os.Exit(1) 
      }

      lineoutput := bufio.NewScanner(file)
      for lineoutput.Scan() {
//	if strings.Contains(lineoutput.Text(), d) {
//            newstr := strings.Replace(lineoutput.Text(), d,  csvd, -1)
//            fmt.Println(newstr)
        fmt.Println(lineoutput.Text())
//        }
      }
      file.Close()
    }
}    

