package main 

import (
    "fmt"
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
    jsonfile, err := os.Open("unixmeta.json")
    if err != nil {
	fmt.Fprintf(os.Stdout, "Unix metdata json file not available or named incorrectly")
        os.Exit(1) 
    }

    // Read json file into a byte array
    b, err := ioutil.ReadAll(jsonfile)
    if err != nil {
	fmt.Fprintf(os.Stdout, "Can not read json file into a byte array")
        os.Exit(1) 
    }

    
    var files Unixfiles 
    json.Unmarshal(b, &files)

    for _,value := range files.Files {
      // Read file line by then
      file, err := os.Open(value.Input)
      if err != nil {
	fmt.Fprintf(os.Stdout, "Unix metdata json file not available or named incorrectly")
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
