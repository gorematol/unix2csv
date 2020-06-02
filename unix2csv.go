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
    jsonfile, err := openfile("unixmeta.json")
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

    var files Unixfiles 
    json.Unmarshal(b, &files)

    for _,value := range files.Files {
      file, err := openfile(value.Input)
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

    jsonfile.Close()
}

func openfile(filename string) (*os.File, error) {
        f, err := os.Open(filename)
        if err != nil {
		return nil, err 
        }
        return f, nil
}
